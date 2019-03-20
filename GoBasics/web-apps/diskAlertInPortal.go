package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)
type DiskSpaceNotificationStruct struct {
	Cid			string	`json:"cid"`
	NotificationState	string	`json:"notificatioState"`
	MaxDiskSpace		int64	`json:"maxDiskSpace"`
	DiskUsed		int64	`json:"diskUsed"`
	FreeDiskSpace		int64	`json:"freeDiskSpace"`
	DiskUsedPercent		int64	`json:"diskUsedPercent"`
	LocationDetails		Location	`json:"locationDetails"`
}
type Location struct {
	LocationID	string
	LocationFQDN	string
}
type MonitorData struct {
	AgentId string `json:"agentId"`
	CPUusage string `json:"CPUusage"`
	DiskUsage DiskStatus `json:"DiskUsage"`
}
type DiskStatus struct {
	All  int64 `json:"all"`
	Used int64 `json:"used"`
	Free int64 `json:"free"`
	UsedPercent string `json:"usedPercent"`
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/diskalert",diskAlertPortalHandler)
	mux.HandleFunc("/metrics",hostMetricsHandler)
	err := http.ListenAndServe("localhost:8282",mux)
	if err != nil{
		fmt.Println("error while starting go server...ERROR:",err)
	}
}
func diskAlertPortalHandler(w http.ResponseWriter, r *http.Request){
	resp := DiskSpaceNotificationStruct{}
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil{
		fmt.Println("error while decoding the reqeust body..ERROR:",err)
	}else {
		fmt.Println("body:\n",resp)
	}
	return
}
func hostMetricsHandler(w http.ResponseWriter, r *http.Request){
	resp := MonitorData{}
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil{
		fmt.Println("error while decoding the reqeust body..ERROR:",err)
	}else {
		fmt.Println("body:\n",resp)
	}
	return
}