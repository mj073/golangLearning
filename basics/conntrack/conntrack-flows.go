package main

import (
	"github.com/vishvananda/netlink/nl"
	"net"
	"fmt"
	"golang.org/x/sys/unix"
	"encoding/binary"
	"bytes"
	"github.com/vishvananda/netns"
	"log"
	"time"
)

const (
	// For Parsing Mark
	TCP_PROTO = 6
	UDP_PROTO = 17
)
var L4ProtoMap = map[uint8]string{
	6:  "tcp",
	17: "udp",
}
type ConntrackTableType uint8
const (
	ConntrackTable = ConntrackTableType(1)
	ConntrackExpectTable = ConntrackTableType(2)
)

// InetFamily Family type
type InetFamily uint8

// The full conntrack flow structure is very complicated and can be found in the file:
// http://git.netfilter.org/libnetfilter_conntrack/tree/include/internal/object.h
// For the time being, the structure below allows to parse and extract the base information of a flow
type ipTuple struct {
	Bytes    uint64
	DstIP    net.IP
	DstPort  uint16
	Packets  uint64
	Protocol uint8
	SrcIP    net.IP
	SrcPort  uint16
}

type ConntrackFlow struct {
	FamilyType uint8
	Forward    ipTuple
	Reverse    ipTuple
	Mark       uint32
}
type Handle struct {
	sockets      map[int]*nl.SocketHandle
	lookupByDump bool
}
func (s *ConntrackFlow) String() string {
	// conntrack cmd output:
	// udp      17 src=127.0.0.1 dst=127.0.0.1 sport=4001 dport=1234 packets=5 bytes=532 [UNREPLIED] src=127.0.0.1 dst=127.0.0.1 sport=1234 dport=4001 packets=10 bytes=1078 mark=0
	return fmt.Sprintf("%s\t%d src=%s dst=%s sport=%d dport=%d packets=%d bytes=%d\tsrc=%s dst=%s sport=%d dport=%d packets=%d bytes=%d mark=%d",
		L4ProtoMap[s.Forward.Protocol], s.Forward.Protocol,
		s.Forward.SrcIP.String(), s.Forward.DstIP.String(), s.Forward.SrcPort, s.Forward.DstPort, s.Forward.Packets, s.Forward.Bytes,
		s.Reverse.SrcIP.String(), s.Reverse.DstIP.String(), s.Reverse.SrcPort, s.Reverse.DstPort, s.Reverse.Packets, s.Reverse.Bytes,
		s.Mark)
}
func (h *Handle) newNetlinkRequest(proto, flags int) *nl.NetlinkRequest {
	// Do this so that package API still use nl package variable nextSeqNr
	if h.sockets == nil {
		return nl.NewNetlinkRequest(proto, flags)
	}
	return &nl.NetlinkRequest{
		NlMsghdr: unix.NlMsghdr{
			Len:   uint32(unix.SizeofNlMsghdr),
			Type:  uint16(proto),
			Flags: unix.NLM_F_REQUEST | uint16(flags),
		},
		Sockets: h.sockets,
	}
}
func (h *Handle) ConntrackTableList(table ConntrackTableType, family InetFamily) ([]*ConntrackFlow, error) {
	res, err := h.dumpConntrackTable(table, family)
	if err != nil {
		return nil, err
	}

	// Deserialize all the flows
	var result []*ConntrackFlow
	for _, dataRaw := range res {
		result = append(result, parseRawData(dataRaw))
	}

	return result, nil
}
func (h *Handle) ConntrackTableFlush(table ConntrackTableType) error {
	req := h.newConntrackRequest(table, unix.AF_INET, nl.IPCTNL_MSG_CT_DELETE, unix.NLM_F_ACK)
	_, err := req.Execute(unix.NETLINK_NETFILTER, 0)
	return err
}
func (h *Handle) newConntrackRequest(table ConntrackTableType, family InetFamily, operation, flags int) *nl.NetlinkRequest {
	// Create the Netlink request object
	req := h.newNetlinkRequest((int(table)<<8)|operation, flags)
	// Add the netfilter header
	msg := &nl.Nfgenmsg{
		NfgenFamily: uint8(family),
		Version:     nl.NFNETLINK_V0,
		ResId:       0,
	}
	req.AddData(msg)
	return req
}

func (h *Handle) dumpConntrackTable(table ConntrackTableType, family InetFamily) ([][]byte, error) {
	req := h.newConntrackRequest(table, family, nl.IPCTNL_MSG_CT_GET, unix.NLM_F_DUMP)
	return req.Execute(unix.NETLINK_NETFILTER, 0)
}
const (
	// backward compatibility with golang 1.6 which does not have io.SeekCurrent
	seekCurrent = 1
)
func parseRawData(data []byte) *ConntrackFlow {
	s := &ConntrackFlow{}
	var proto uint8
	// First there is the Nfgenmsg header
	// consume only the family field
	reader := bytes.NewReader(data)
	binary.Read(reader, nl.NativeEndian(), &s.FamilyType)

	// skip rest of the Netfilter header
	reader.Seek(3, seekCurrent)
	// The message structure is the following:
	// <len, NLA_F_NESTED|CTA_TUPLE_ORIG> 4 bytes
	// <len, NLA_F_NESTED|CTA_TUPLE_IP> 4 bytes
	// flow information of the forward flow
	// <len, NLA_F_NESTED|CTA_TUPLE_REPLY> 4 bytes
	// <len, NLA_F_NESTED|CTA_TUPLE_IP> 4 bytes
	// flow information of the reverse flow
	for reader.Len() > 0 {
		if nested, t, l := parseNfAttrTL(reader); nested {
			switch t {
			case nl.CTA_TUPLE_ORIG:
				if nested, t, _ = parseNfAttrTL(reader); nested && t == nl.CTA_TUPLE_IP {
					proto = parseIpTuple(reader, &s.Forward)
				}
			case nl.CTA_TUPLE_REPLY:
				if nested, t, _ = parseNfAttrTL(reader); nested && t == nl.CTA_TUPLE_IP {
					parseIpTuple(reader, &s.Reverse)
				} else {
					// Header not recognized skip it
					reader.Seek(int64(l), seekCurrent)
				}
			case nl.CTA_COUNTERS_ORIG:
				s.Forward.Bytes, s.Forward.Packets = parseByteAndPacketCounters(reader)
			case nl.CTA_COUNTERS_REPLY:
				s.Reverse.Bytes, s.Reverse.Packets = parseByteAndPacketCounters(reader)
			}
		}
	}
	if proto == TCP_PROTO {
		reader.Seek(64, seekCurrent)
		_, t, _, v := parseNfAttrTLV(reader)
		if t == nl.CTA_MARK {
			s.Mark = uint32(v[3])
		}
	} else if proto == UDP_PROTO {
		reader.Seek(16, seekCurrent)
		_, t, _, v := parseNfAttrTLV(reader)
		if t == nl.CTA_MARK {
			s.Mark = uint32(v[3])
		}
	}
	return s
}
func parseNfAttrTLV(r *bytes.Reader) (isNested bool, attrType, len uint16, value []byte) {
	isNested, attrType, len = parseNfAttrTL(r)

	value = make([]byte, len)
	binary.Read(r, binary.BigEndian, &value)
	return isNested, attrType, len, value
}
func parseNfAttrTL(r *bytes.Reader) (isNested bool, attrType, len uint16) {
	binary.Read(r, nl.NativeEndian(), &len)
	len -= nl.SizeofNfattr

	binary.Read(r, nl.NativeEndian(), &attrType)
	isNested = (attrType & nl.NLA_F_NESTED) == nl.NLA_F_NESTED
	attrType = attrType & (nl.NLA_F_NESTED - 1)

	return isNested, attrType, len
}
// This method parse the ip tuple structure
// The message structure is the following:
// <len, [CTA_IP_V4_SRC|CTA_IP_V6_SRC], 16 bytes for the IP>
// <len, [CTA_IP_V4_DST|CTA_IP_V6_DST], 16 bytes for the IP>
// <len, NLA_F_NESTED|nl.CTA_TUPLE_PROTO, 1 byte for the protocol, 3 bytes of padding>
// <len, CTA_PROTO_SRC_PORT, 2 bytes for the source port, 2 bytes of padding>
// <len, CTA_PROTO_DST_PORT, 2 bytes for the source port, 2 bytes of padding>
func parseIpTuple(reader *bytes.Reader, tpl *ipTuple) uint8 {
	for i := 0; i < 2; i++ {
		_, t, _, v := parseNfAttrTLV(reader)
		switch t {
		case nl.CTA_IP_V4_SRC, nl.CTA_IP_V6_SRC:
			tpl.SrcIP = v
		case nl.CTA_IP_V4_DST, nl.CTA_IP_V6_DST:
			tpl.DstIP = v
		}
	}
	// Skip the next 4 bytes  nl.NLA_F_NESTED|nl.CTA_TUPLE_PROTO
	reader.Seek(4, seekCurrent)
	_, t, _, v := parseNfAttrTLV(reader)
	if t == nl.CTA_PROTO_NUM {
		tpl.Protocol = uint8(v[0])
	}
	// Skip some padding 3 bytes
	reader.Seek(3, seekCurrent)
	for i := 0; i < 2; i++ {
		_, t, _ := parseNfAttrTL(reader)
		switch t {
		case nl.CTA_PROTO_SRC_PORT:
			parseBERaw16(reader, &tpl.SrcPort)
		case nl.CTA_PROTO_DST_PORT:
			parseBERaw16(reader, &tpl.DstPort)
		}
		// Skip some padding 2 byte
		reader.Seek(2, seekCurrent)
	}
	return tpl.Protocol
}
func parseBERaw16(r *bytes.Reader, v *uint16) {
	binary.Read(r, binary.BigEndian, v)
}

func parseBERaw64(r *bytes.Reader, v *uint64) {
	binary.Read(r, binary.BigEndian, v)
}
func parseByteAndPacketCounters(r *bytes.Reader) (bytes, packets uint64) {
	for i := 0; i < 2; i++ {
		switch _, t, _ := parseNfAttrTL(r); t {
		case nl.CTA_COUNTERS_BYTES:
			parseBERaw64(r, &bytes)
		case nl.CTA_COUNTERS_PACKETS:
			parseBERaw64(r, &packets)
		default:
			return
		}
	}
	return
}
func NewHandle(nlFamilies ...int) (*Handle, error) {
	return newHandle(netns.None(), netns.None(), nlFamilies...)
}
func newHandle(newNs, curNs netns.NsHandle, nlFamilies ...int) (*Handle, error) {
	h := &Handle{sockets: map[int]*nl.SocketHandle{}}
	fams := nl.SupportedNlFamilies
	if len(nlFamilies) != 0 {
		fams = nlFamilies
	}
	for _, f := range fams {
		s, err := nl.GetNetlinkSocketAt(newNs, curNs, f)
		if err != nil {
			return nil, err
		}
		h.sockets[f] = &nl.SocketHandle{Socket: s}
	}
	return h, nil
}
func main(){
	h,err := NewHandle(unix.NETLINK_NETFILTER)
	if err != nil {
		log.Fatalln("failed to create Handle..ERROR:",err)
	}
	fmt.Println("handle created...")
	err = h.ConntrackTableFlush(ConntrackTable)
	if err != nil {
		log.Fatalln("failed to flush conntrack table..ERROR:", err)
	}
	for {
		err = h.ConntrackTableFlush(ConntrackTable)
		if err != nil {
			log.Fatalln("failed to flush conntrack table..ERROR:", err)
		}
		flows, err := h.ConntrackTableList(ConntrackTable, InetFamily(unix.AF_INET))
		if err != nil {
			log.Fatalln("failed to get conntrack flow..ERROR:", err)
		}
		if len(flows) == 0 {
			fmt.Println("no flows")
		}
		for _, flow := range flows {
			fmt.Println(flow)
		}
		<-time.After(time.Millisecond * 50)
	}
}