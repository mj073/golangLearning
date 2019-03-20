package main

import (
	"net"
	"fmt"
)

func udpFlowCreateProg(flows, srcPort int, dstIP string, dstPort int) {
	for i := 0; i < flows; i++ {
		ServerAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", dstIP, dstPort))
		if err != nil {
			fmt.Errorf("%s","failed to resolve dst udp addr..ERROR:"+err.Error())
		}

		LocalAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", srcPort+i))
		if err != nil {
			fmt.Errorf("%s","failed to resolve src udp addr..ERROR:"+err.Error())
		}

		Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
		if err != nil {
			fmt.Errorf("%s","failed DialUDP..ERROR:"+err.Error())
		}
		Conn.Write([]byte("Hello World"))
		Conn.Close()
	}
}
func main(){
	udpFlowCreateProg(5,5000,"127.0.0.1",6000)
}

