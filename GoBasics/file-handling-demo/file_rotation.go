package main

import (
	"os"
	"path/filepath"
	"time"
	"log"
	"strings"
	"flag"
	"os/exec"
)
type FileStruct struct{
	Filename string
	CreatedAt time.Time
	FileInfo	os.FileInfo
}
type FileRotate struct {
	Dir string
	File chan *FileStruct
}

const(
	MAX_FILE_SIZE  = 1024
	MAX_FILE_DURATION = time.Minute * 1
	filename_time_format = "20060102150405000"
	MAX_TRY = 5
)
var blockingChan chan int
var globalFileChan chan FileStruct
func main(){
	path := flag.String("dir", "", "absolute path of dir ")
	flag.Parse()

	if strings.Contains(*path, "./") {
		log.Fatalln("ERROR: please give absolute path")
	}
	if info, err := os.Stat(*path); err == nil{
		if ! info.IsDir(){
			log.Fatalln(*path," is not a directory")
		}
		log.Println("directory found..")
	} else {
		if os.IsNotExist(err){
			log.Println("directory not found..")
			log.Println("creating the directory..",*path)
			if err := exec.Command("mkdir","-p",*path).Run(); err != nil{
				log.Fatalln("failed to create the directory ERROR:",err)
			}
			log.Println("directory created successfully")
		}
	}
	filerotate := &FileRotate{*path,make(chan *FileStruct,1)}
	log.Println("starting file operations routine...")
	go filerotate.FileOperationsRoutine()

	//go filerotate.MonitorFile()

	log.Println("generating file name struct..")
	filerotate.File <- GetFileStruct()
	<- blockingChan
}

func (rotate *FileRotate) FileOperationsRoutine(){
	try := 0
	var f *os.File
	//var err error
	for{
		//<- time.After(time.Millisecond * 50) // with the help of this we can optimize CPU utilization
		                                      // strace -c ./bin/file_rotation -dir "/tmp/file_rotaion" (try to run with and without this statement to see effect)
		if file, ok := <- rotate.File; ok{
			if f == nil {
				log.Println("WARN: file ptr is nil")
			}
			filePath := filepath.Join(rotate.Dir, file.Filename)
			//rotate.CreateFile(filePath)
			//if file.FileInfo == nil{
			//	log.Println("file.FileInfo is nil")
			//	file.FileInfo, err = os.Stat(filePath)
			//	if err != nil && os.IsNotExist(err) {
			//		log.Println("file:", filePath, " does not exist...creating file")
			//		_, err = os.Create(filePath)
			//		if err != nil {
			//			log.Println("failed to create the file ERROR:",err)
			//			try++
			//			if try == MAX_TRY {
			//				log.Println("tried creating the file ",MAX_TRY," times. No luck")
			//				time.Sleep(time.Second * 3)
			//				continue
			//			}
			//			rotate.File <- file
			//			continue
			//		}
			//		log.Println("file:", filePath, " created successfully")
			//		//fileInfo,err = os.Stat(filePath)
			//		file.FileInfo, _ = os.Stat(filePath)
			//	}
			//}
			fileInfo, err := os.Stat(filePath)
			if err != nil && os.IsNotExist(err) {
				log.Println("file:", filePath, " does not exist...creating file")
				_, err = os.Create(filePath)
				if err != nil {
					log.Println("failed to create the file ERROR:",err)
					try++
					if try == MAX_TRY {
						log.Println("tried creating the file ",MAX_TRY," times. No luck")
						time.Sleep(time.Second * 3)
						continue
					}
					rotate.File <- file
					continue
				}
				log.Println("file:", filePath, " created successfully")
				fileInfo,err = os.Stat(filePath)
			}
			//fileInfo, err := os.Stat(filePath)

			//err := rotate.CreateFile(filePath)
			//if err != nil {
			//
			//}
			sizeCheck := fileInfo.Size() >= MAX_FILE_SIZE
			durationCheck := time.Now().After(file.CreatedAt.Add(MAX_FILE_DURATION))
			if sizeCheck || durationCheck {
				log.Println("filesize of ",filePath," is ",fileInfo.Size(),"..filesizeCheck=",sizeCheck,"...fileDurationCheck=",durationCheck)
				log.Println("rotating the file..")
				f.Close()
				f = nil
				go ZipAndSendRoutine(filePath)
				rotate.File <- GetFileStruct()
			}else{
				if f == nil {
					f, err = os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0644)
					if err != nil {
						log.Println("failed to open the file ERROR:", err)
						try++
						if try == MAX_TRY {
							log.Println("tried opening the file ", MAX_TRY, " times. No luck")
							time.Sleep(time.Second * 3)
							continue
						}
						rotate.File <- file
						continue
					}
					log.Println("file opened in append mode")
				}
				rotate.File <- file
			}
		}
	}
	log.Println("***FileOperations go routine died abruptly***")
}
//func (rotate *FileRotate) CreateFile(filePath string) (Err error){
//	_, err := os.Stat(filePath)
//	try := 0
//	if err != nil && os.IsNotExist(err) {
//		log.Println("file:", filePath, " does not exist...creating file")
//		f, Err := os.Create(filePath)
//		defer f.Close()
//		if Err != nil {
//			log.Println("failed to create the file ERROR:", Err)
//			try++
//			if try == MAX_TRY {
//				log.Println("tried creating the file ", MAX_TRY, " times. No luck")
//				time.Sleep(time.Second * 3)
//				return Err
//			}
//		}
//		log.Println("file:", filePath, " created successfully")
//	}
//	return Err
//}
//
//func (rotate *FileRotate) MonitorFile(){
//	var file FileStruct
// 	for {
//		if file,ok := <- globalFileChan;ok{
//			filePath := filepath.Join(rotate.Dir, file.Filename)
//			fileInfo, _ := os.Stat(filePath)
//			sizeCheck := fileInfo.Size() >= MAX_FILE_SIZE
//			durationCheck := time.Now().After(file.CreatedAt.Add(MAX_FILE_DURATION))
//			if sizeCheck || durationCheck {
//				log.Println("filesize of ",filePath," is ",fileInfo.Size(),"..filesizeCheck=",sizeCheck,"...fileDurationCheck=",durationCheck)
//				log.Println("rotating the file..")
//				io
//				f.Close()
//				f = nil
//				go ZipAndSendRoutine(filePath)
//				rotate.File <- GetFileStruct()
//			}
//		}
//
//	}
//}

func GetFileStruct() *FileStruct{
	current_time := time.Now()
	log.Println("returning the filestruct..")
	return &FileStruct{"example_" + current_time.Format(filename_time_format),current_time,nil}
}

func ZipAndSendRoutine(file string){
	log.Println("zipping and sending the file:",file,"to remote server")
}