package main

import (
	"flag"
	"net/http"
	"os"
	"log"
	"strings"
)
var f *os.File
func main() {
	r := http.NewServeMux()
	logPath := flag.String("logfile","","Log file absolute path")
	flag.Parse()
	if strings.Contains(*logPath,"./") {
		log.Fatalln("please provide absolute log path")
	}
	if fileInfo,err := os.Stat(*logPath); os.IsNotExist(err){
		log.Println("Creating the log file ",fileInfo.Name())
		if f,err = os.Create(*logPath); err != nil{
			log.Fatalln("failed to create log file... ERROR:",err)
		}
	}
	log.SetOutput(f)
	if f != nil{
		r.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {log.Println("web root is requested..")})
	}else{
		log.Fatalln("failed to start the server..ERROR: log file pointer is nil")
	}

	http.ListenAndServe(":9000",r)
}
