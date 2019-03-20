package main

import (
	"log"
	"os/exec"
	"flag"
	"os"
	"strings"
)

func main(){
	file := flag.String("file","","absolute filepath to zip")
	server := flag.String("server","","sftp server")
	user := flag.String("user","","username of dest server")
	pswd := flag.String("password","","pswd of dest sftp server")
	flag.Parse()
	if strings.Contains(*file,"./"){
		log.Fatalln("please use absolute filepath..")
	}

	fileTemp := *file + "_temp"
	fileBkp := *file+"_bkp"
	zipFile := *file+".gz"
	//ftp_cmd := "ftp -in "+*server+" << END\n" + "user "+*user+" "+*pswd+"\n"+"binary\n"+"put "+zipFile+"\n"+"END"
	scp_cmd := "scp "+zipFile + user+"@"+server+":/home/"+user+"/"
	serial_no_cmd := "awk '{print NR\",\"$0}' "+ *file + "> " + fileTemp + " && mv " + *file + " " + fileBkp + " && mv " + fileTemp + " " + *file + " && rm " + fileBkp
	log.Println("cmd: ",serial_no_cmd)
	log.Println("adding serial numbers to the file")
	if err := exec.Command("bash","-c",serial_no_cmd).Run(); err != nil{
		log.Fatalln("failed to add serial numbers to the file..ERROR:",err)
	}
	log.Println("serial numbers added successfully..")
	log.Println("now trying to zip the file..")
	if err := exec.Command("gzip",*file).Run(); err != nil{
		log.Fatalln("failed to zip the file ERROR:",err)
	}

	fileInfo,err := os.Stat(zipFile)
	if err != nil {
		if os.IsNotExist(err){
			log.Fatalln("zipped file ",zipFile," doesnt exist")
		}else{
			log.Fatalln("error while getting stats on file ERROR:",err)
		}
	}else {
		log.Println("successfully zipped the file")
		if fileInfo.Size() >= 3145728{
			log.Println("filesize threshold check failed")
			log.Println("filesize is greater than threshold value")
		}else{
			log.Println("filesize threshold check success")
			log.Println("transferring the file to destination server..")
		}
		if err := exec.Command("bash","-c",ftp_cmd).Run(); err != nil{
			log.Fatalln("failed to ftp the file ERROR:",err)
		}else {
			log.Println("successfully transferred the file")
		}
	}
}
