package main

import (
	//"fmt"
	//"encoding/binary"
	"fmt"
	"strings"
)

type ContainerStats struct{
	ContainerId	string  	`json:"containerId"`
	Name		string		`json:"name"`
	Status 		string		`json:"status"`
	ImageName 	string		`json:"image"`
	ContentSize 	string		`json:"contentSize"`
	UpSince		string		`json:"upSince"`
	CPUAvgLoad	string		`json:"cpuAvgLoad"`
	PerCPUUsage 	string 		`json:"perCPUusage"`
	TotalCPUUsage	string		`json:"totalCPUusage"`
	UserCPUStat	string		`json:"userCPUstat"`
	SystemCPUStat	string		`json:"sytemCPUstat"`
	MemLimit	string		`json:"memLimit"`
	MemUsage	string		`json:"memUsage"`
	MemUsagePercent	string		`json:"memUsagePercent"`
	NetworkStat	NetStatMeta	`json:"networkStat"`
	//IOStat		IOStatMeta	`json:"ioStat"`
}

type NetStatMeta struct{
	Iface		string
	RecvBytes	string
	RecvPckts	string
	RecvErrs	string
	RecvDrop	string
	RecvFIFO	string
	RecvFrame	string
	RecvCompressed	string
	RecvMulticast	string
	TransBytes	string
	TransPckts	string
	TransErrs	string
	TransDrop	string
	TransFIFO	string
	TransFrame	string
	TransCompressed	string
	TransMulticast	string
}

type IOStatMeta struct{
	Read	string
	Write	string
	Sync	string
	Async	string
	Total	string
}

func main(){
	part := "6f2ae7fd28e2","lr1/cdn_lr001_m1_clrtrp","running","alef-cdnv2-child:1.2","165166","2016-07-04T22:30:41","0.55,0.42,0.37","40303587590","40303587590","1486","2686","18446744073709551615","10993664","1.05%",{"eth0","4432555","33596","0","0","0","0","0","0","4094940","43950","0","0","0","0","0","0"},{"26128384","0","0","26128384","26128384"}
	//fields := strings.Split(part," ")
	//fmt.Println(fields[0:14])
	//netMeta := NetStatMeta{fmt.Sprint(fields[14:31])}
	//fmt.Println("NetMeta:",netMeta)
	//fmt.Println(fields[31:36],"...",len(fields))
	//container := ContainerStats{"6f2ae7fd28e2","lr1/cdn_lr001_m1_clrtrp","running","alef-cdnv2-child:1.2","165166","2016-07-04T07:14:23","0.55,0.34,0.32","39066501292","39066501292","1418","2630","18446744073709551615","12025856","1.15%","eth0","4345017","32909","0","0","0","0","0","0","4014546","43065","0","0","0","0","0","0","26001408","0","0","26001408","26001408"}
	//container := ContainerStats{"6f2ae7fd28e2","lr1/cdn_lr001_m1_clrtrp","running","alef-cdnv2-child:1.2","165166","2016-07-04T07:14:23","0.55,0.34,0.32","39066501292","39066501292","1418","2630","18446744073709551615","12025856","1.15%",NetStatMeta{"eth0","4345017","32909","0","0","0","0","0","0","4014546","43065","0","0","0","0","0","0"}}
	////fmt.Sprint(fields[:14]),fmt.Sprint(fields[14:31]),fmt.Sprint(fields[31:])
	fmt.Println(container)
	strings.

}
