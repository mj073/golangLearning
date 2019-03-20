package lib

import (
	"github.com/google/gopacket/routing"
	"fmt"
	"time"
	"net"
)
var RoutingTable routing.Router
func UpdateRoutingTableInfo(){
	for {
		r,err := routing.New()
		if err != nil {
			fmt.Errorf("%s","failed to update routing table..ERROR: "+err.Error())
		}
		fmt.Println("routing table r :",r)
		RoutingTable = r
		fmt.Println("routing table RoutingTable:",RoutingTable)
		<- time.After(time.Minute * 1)
	}
}
func GetInterface(src,dst string) string{
	if RoutingTable != nil {
		iff := &net.Interface{}
		var err error
		if src == "" {
			iff, _, _, err = RoutingTable.Route(net.ParseIP(dst))
		}else {
			iff, _, _, err = RoutingTable.RouteWithSrc(nil,net.ParseIP(src),net.ParseIP(dst))
		}
		if err == nil && iff != nil{
			return iff.Name
		}else {
			fmt.Println("interface is nil..ERROR:",err)
		}
	}
	return "NA"
}
