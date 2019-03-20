package main

import (
	"syscall"
	"fmt"
	"io"
	"unsafe"
	"time"
)
//type Message struct{
//	Hdr syscall.NlMsghdr
//	Gen syscall.IfInfomsg
//}
type Hdr syscall.NlMsghdr
type IfInfoMsg syscall.IfInfomsg

func main(){
	fd, err := syscall.Socket(syscall.AF_NETLINK,syscall.SOCK_RAW,syscall.NETLINK_ROUTE)
	if err != nil{
		panic(err)
	}
	defer syscall.Close(fd)
	fmt.Println("socket fd:",fd)
	sa := syscall.SockaddrNetlink{Family: syscall.AF_NETLINK,}
	err = syscall.Bind(fd,&sa)
	if err != nil{
		panic(err)
	}
	//msgtype := syscall.RTM_GETLINK
	//msgtype := syscall.RTM_GETROUTE
	nlreq,err := NewMessage(
		Hdr{Type: syscall.RTM_GETROUTE, Pid: uint32(syscall.Getpid()), Flags:syscall.NLM_F_REQUEST | syscall.NLM_F_DUMP, },
		IfInfoMsg{Family:syscall.AF_PACKET},
	)
	//iov := &syscall.Iovec{}
	//iov.Base = []byte(fmt.Sprintf("%v",nlreq))
	//iov.Len = nlreq.Hdr.Len

 	//msghdr := &syscall.Msghdr{}
	//msghdr.Name = &byte(sa)
	//msghdr.Namelen = len(sa)
	//msghdr.Iov = &iov
	//msghdr.Iovlen = 1
	//epfd, e3 := syscall.EpollCreate1(0)
	//if e3 != nil {
	//	panic(e3)
	//}
	//ev := syscall.EpollEvent{ Events: syscall.EPOLLIN, Fd: int32(fd) }
	//e4 := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &ev)
	//if e4 != nil {
	//	panic(e4)
	//}
	//defer func(){
	//	syscall.EpollCtl(epfd, syscall.EPOLL_CTL_DEL, fd, &ev)
	//}()
	fmt.Println("waiting to receive msg from kernel....")
	go gorx(fd,0)
	<- time.After(time.Second * 1)
	fmt.Println("after 1 sec sleep")
	fmt.Println("netlink request:",nlreq)
	err = syscall.Sendto(fd,nlreq,0,&sa)
	if err != nil{
		fmt.Println("failed to send msg req to kernel...ERROR:",err)
		return
	}

	<- time.After(time.Second * 1000)
}


// convert msg to []byte
func NewMessage(hdr Hdr,msg io.Reader) ([]byte, error){
	b := make([]byte, syscall.Getpagesize())
	nmsg, err := msg.Read(b[syscall.NLMSG_HDRLEN:])
	if err != nil {
		return nil, err
	}
	//Align message hdrs and msg
	n := func(i int) int{
		return (i + int(syscall.NLMSG_ALIGNTO) - 1) & ^(int(syscall.NLMSG_ALIGNTO) - 1)
	}(syscall.NLMSG_HDRLEN + nmsg)
	//na, err := ReadAllAttrs(b[n:], attrs...)
	//if err != nil {
	//	return nil, err
	//}
	//n += na
	hdr.Len = uint32(n)
	if _,err := hdr.Read(b);err != nil{
		return nil, err
	}
	return b[:n], nil
}

func (msg IfInfoMsg) Read(b []byte) (int,error){
	*(*IfInfoMsg)(unsafe.Pointer(&b[0])) = msg
	return syscall.SizeofIfInfomsg, nil
}
func (hdr *Hdr) Read(b []byte) (int,error){
	*HdrPtr(b) = *hdr
	return syscall.NLMSG_HDRLEN, nil
}
func HdrPtr(b []byte) *Hdr {
	if len(b) < syscall.NLMSG_HDRLEN {
		return nil
	}
	return (*Hdr)(unsafe.Pointer(&b[0]))
}

func gorx(fd,epfd int){
	rb := make([]byte, syscall.Getpagesize())
	//evs := make([]syscall.EpollEvent, 1)
	for {
		//if n, err := syscall.EpollWait(epfd, evs, 10); err!=nil {
		//	panic(err)
		//}else if n > 0 {
			nr, _, err := syscall.Recvfrom(fd, rb, 0)
			if err != nil {
				continue
			}
			if nr > 0 {
				if nr < syscall.NLMSG_HDRLEN {
					fmt.Errorf("Got short response from netlink")
					continue
				}
				//rb = rb[:nr]
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
							case syscall.RTA_GATEWAY:
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
					case syscall.RTM_NEWROUTE:
						fmt.Println("New Route msg received from kernel...")
						attrs, err := syscall.ParseNetlinkRouteAttr(&msg)
						if err != nil {
							fmt.Println("error while parsing netlink route attributes..ERROR:", err)
							continue
						}
						for _, atr := range attrs {
							fmt.Println("attr Type:", atr.Attr.Type)
							fmt.Println("atr.Value:",atr.Value)
							//switch atr.Attr.Type {
							////case syscall.IFLA_ADDRESS:
							////	fmt.Println("attribute Interface Link Address:", atr.Value)
							////case syscall.IFLA_LINK:
							////	fmt.Println("attribute Interface Link Type:", atr.Value)
							//case syscall.IFLA_IFNAME:
							//	fmt.Println("attribute Interface Link Name:", string(atr.Value))
							//case syscall.RTA_GATEWAY:
							//	fmt.Println("attribute Interface Link Name:", string(atr.Value))
							//default:
							//	fmt.Println("None of attributes matched")
							//	continue
							//}
						}
					default :
						fmt.Println("none of the msg types..")
						continue loop
					}
				}
			} else {
				fmt.Println("nothing received yet...")
			}
		//}
	}
}