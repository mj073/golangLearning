package main

import (
	"fmt"
	"regexp"
	"strings"
	"bufio"
)

func main() {
	data := `20: xeth16: <no-carrier,broadcast,multicast,up> mtu 1500 qdisc noqueue state lower-down mode default group default qlen 1000
    link/ether 50:18:4c:00:11:8b brd ff:ff:ff:ff:ff:ff
    inet 10.0.16.47/24 scope global xeth16
       valid_lft forever preferred_lft forever
22: xeth18: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:8d brd ff:ff:ff:ff:ff:ff
    inet 10.0.18.47/24 scope global xeth18
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:118d/64 scope link
       valid_lft forever preferred_lft forever
24: xeth20: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:8f brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.47/24 scope global xeth20
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:118f/64 scope link
       valid_lft forever preferred_lft forever
26: xeth22: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:91 brd ff:ff:ff:ff:ff:ff
    inet 10.0.22.47/24 scope global xeth22
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:1191/64 scope link
       valid_lft forever preferred_lft forever
28: xeth24: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:93 brd ff:ff:ff:ff:ff:ff
    inet 10.0.24.47/24 scope global xeth24
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:1193/64 scope link
       valid_lft forever preferred_lft forever
30: xeth26: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:95 brd ff:ff:ff:ff:ff:ff
    inet 10.0.26.47/24 scope global xeth26
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:1195/64 scope link
       valid_lft forever preferred_lft forever
32: xeth28: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:97 brd ff:ff:ff:ff:ff:ff
    inet 10.0.28.47/24 scope global xeth28
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:1197/64 scope link
       valid_lft forever preferred_lft forever
34: xeth30: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:99 brd ff:ff:ff:ff:ff:ff
    inet 10.0.30.47/24 scope global xeth30
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:1199/64 scope link
       valid_lft forever preferred_lft forever
36: xeth32: <broadcast,multicast,up,lower-up> mtu 1500 qdisc noqueue state up mode default group default qlen 1000
    link/ether 50:18:4c:00:11:9b brd ff:ff:ff:ff:ff:ff
    inet 10.0.32.47/24 scope global xeth32
       valid_lft forever preferred_lft forever
    inet6 fe80::5218:4cff:fe00:119b/64 scope link
       valid_lft forever preferred_lft forever
37: docker0: <no-carrier,broadcast,multicast,up> mtu 1500 qdisc noqueue state down mode default group default qlen default
    link/ether 02:42:f2:12:e9:6d brd ff:ff:ff:ff:ff:ff
    inet 172.26.0.1/16 scope global docker0
       valid_lft forever preferred_lft forever
	`
	re := regexp.MustCompile(`inet (.*) scope global (.*)`)
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		s := re.FindStringSubmatch(scanner.Text())
		if len(s) != 0{
			fmt.Println(s[2], "\t", s[1])
		}
	}


}
