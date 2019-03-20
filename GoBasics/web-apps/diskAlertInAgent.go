package main

import (
	"net/http"
	"fmt"
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

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/diskalert",diskAlertAgentHandler)
	err := http.ListenAndServe(":8080",mux)
	if err != nil{
		fmt.Println("error while starting go server...ERROR:",err)
	}
}
func diskAlertAgentHandler(w http.ResponseWriter, r *http.Request){
	//resp := DiskSpaceNotificationStruct{}
	//err := json.NewDecoder(r.Body).Decode(&resp)
	//if err != nil{
	//	fmt.Println("error while decoding the reqeust body..ERROR:",err)
	//}else {
	//	fmt.Println("body:\n",resp)
	//}
	//return
	resp,err := http.Post("http://192.168.0.25:8181/diskalertManager",r.Header.Get("Content-Type"),r.Body)
	if err != nil{
		fmt.Println("error while calling manager's disk alert API...ERROR:",err)
	}else if resp != nil && resp.StatusCode == http.StatusOK {
		fmt.Println("successfully sent disk alert info to manager...")
	}else{
		fmt.Println("failed to send disk alert info to manager...")
	}
}