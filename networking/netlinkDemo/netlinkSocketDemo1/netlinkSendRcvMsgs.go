package main

import (
	"syscall"
	"fmt"
	"time"
	"unsafe"
	"os"
)
type NetlinkRequestData interface {
	Len() int
	Serialize() []byte
}
type NetlinkRequest struct {
	syscall.NlMsghdr
	Data []NetlinkRequestData
}
func (req *NetlinkRequest) Serialize() []byte {
	length := syscall.SizeofNlMsghdr
	dataBytes := make([][]byte, len(req.Data))
	for i, data := range req.Data {
		dataBytes[i] = data.Serialize()
		length = length + len(dataBytes[i])
	}
	req.Len = uint32(length)
	b := make([]byte, length)
	hdr := (*(*[syscall.SizeofNlMsghdr]byte)(unsafe.Pointer(req)))[:]
	next := syscall.SizeofNlMsghdr
	copy(b[0:next], hdr)
	for _, data := range dataBytes {
		for _, dataByte := range data {
			b[next] = dataByte
			next = next + 1
		}
	}
	return b
}

func (req *NetlinkRequest) AddData(data NetlinkRequestData) {
	if data != nil {
		req.Data = append(req.Data, data)
	}
}
func NewNetlinkRequest() *NetlinkRequest {
	return &NetlinkRequest{
		NlMsghdr: syscall.NlMsghdr{
			Len:   uint32(0),
			Type:  uint16(syscall.NLMSG_DONE),
			Flags: uint16(0),
			Seq:   uint32(0),
			Pid:   uint32(os.Getpid()),
		},
	}
}

func main(){
	fd, err := syscall.Socket(syscall.AF_NETLINK,syscall.SOCK_RAW,syscall.NETLINK_ROUTE)
	if err != nil{
		panic(err)
	}
	defer syscall.Close(fd)
	fmt.Println("socket fd:",fd)
	sa := syscall.SockaddrNetlink{Family: syscall.AF_NETLINK,}
	sa.Pid = 0
	sa.Groups = syscall.RTNLGRP_LINK | syscall.RTNLGRP_IPV4_IFADDR |syscall.RTNLGRP_NOTIFY
	err = syscall.Bind(fd,&sa)
	if err != nil{
		panic(err)
	}
	epfd, e3 := syscall.EpollCreate1(0)
	if e3 != nil {
		panic(e3)
	}
	ev := syscall.EpollEvent{ Events: syscall.EPOLLIN, Fd: int32(fd) }
	e4 := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &ev)
	if e4 != nil {
		panic(e4)
	}
	defer func(){
		syscall.EpollCtl(epfd, syscall.EPOLL_CTL_DEL, fd, &ev)
	}()
	go gorx(fd,epfd)
	//go gotx(fd,NewNetlinkRequest())
	<- time.After(time.Second * 1000)
}
func gorx(fd,epfd int){
	rb := make([]byte, syscall.Getpagesize())
	evs := make([]syscall.EpollEvent, 1)
	for {
		if n, err := syscall.EpollWait(epfd, evs, 10); err!=nil {
			panic(err)
		}else if n > 0 {
			nr, _, err := syscall.Recvfrom(fd, rb, 0)
			if err != nil {
				continue
			}
			if nr > 0 {
				if nr < syscall.NLMSG_HDRLEN {
					fmt.Errorf("Got short response from netlink")
					continue
				}
				rb = rb[:nr]
				msgs, err := syscall.ParseNetlinkMessage(rb[:nr])
				if err != nil {
					fmt.Println("error while parsing netlink messages..ERROR:", err)
					continue
				}
				loop:
				for _, msg := range msgs {
					fmt.Println("message recived:", msg)
					switch msg.Header.Type {
					// Refer http://man7.org/linux/man-pages/man7/rtnetlink.7.html for Message types & attributes
					case syscall.NLMSG_DONE:
						fmt.Println("Netlink DONE msg received")
						break loop
					case syscall.RTM_DELLINK:
						fmt.Println("Del Link msg received from kernel...")
					case syscall.RTM_NEWLINK:
						fmt.Println("New Link msg received from kernel...")
						attrs, err := syscall.ParseNetlinkRouteAttr(&msg)
						if err != nil {
							fmt.Println("error while parsing netlink route attributes..ERROR:", err)
							continue
						}
						for _, atr := range attrs {
							//fmt.Println("attr Type:", atr.Attr.Type)
							switch atr.Attr.Type {
							//case syscall.IFLA_ADDRESS:
							//	fmt.Println("attribute Interface Link Address:", atr.Value)
							//case syscall.IFLA_LINK:
							//	fmt.Println("attribute Interface Link Type:", atr.Value)
							case syscall.IFLA_IFNAME:
								fmt.Println("attribute Interface Link Name:", string(atr.Value))
							//default:
							//	fmt.Println("None of attributes matched")
							//	continue
							}
						}
					case syscall.RTM_GETLINK:
						fmt.Println("Get Link msg received from kernel...")
					case syscall.RTM_DELNEIGH:
						fmt.Println("Del Neigh msg received from kernel...")
					case syscall.RTM_NEWNEIGH:
						fmt.Println("New Neigh msg received from kernel...")
					case syscall.RTM_NEWADDR:
						fmt.Println("New address msg received from kernel...")
					default :
						fmt.Println("none of the msg types..")
						continue loop
					}
				}
			} else {
				fmt.Println("nothing received yet...")
			}
		}
	}
}

//func gotx(fd int,request *NetlinkRequest) {
//	request.Data = []byte{}
//	sockaddr := &syscall.SockaddrNetlink{
//		Family: syscall.AF_NETLINK,
//		Pid:    0,
//		Groups: 0,
//	}
//	if err := syscall.Sendto(fd, request.Serialize(), syscall.NLM_F_ECHO, sockaddr); err != nil {
//		return err
//	}
//}
