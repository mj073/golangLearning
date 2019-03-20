package main

import (
	"os"
	"log"
)

func main() {
	imagepath := "E:/golang/src/GoBasics/stegnography/tiger.jpg"
	newImagePath := "E:/golang/src/GoBasics/stegnography/tiger_new.jpg"
	msg := "hello"
	msg_bytes := []byte(msg)
	len_msg := len(msg_bytes)
	log.Printf("msg: %s\nmsg_bytes: %v\nlen_bytes:%v",msg,msg_bytes,len_msg)
	file,err := os.OpenFile(imagepath,os.O_RDONLY,0644)
	if err != nil{
		log.Fatalln("error while opening the file..ERROR:",err)
	}
	log.Println(file.Name())
	defer file.Close()
	fileInfo, err := file.Stat()

	b := make([]byte,fileInfo.Size())
	//buf := bytes.NewBuffer(b)
	numread, err := file.Read(b)
	if err != nil {
		log.Fatalln("error while reading file..ERROR:",err)
	}
	log.Println("bytes read from file: ",numread)

	newImageBytes := createBytes(msg_bytes,b,len_msg,numread)
	_, err = os.Stat(newImagePath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(newImagePath)
		if err != nil {
			log.Fatalln("error while creating the path...ERROR:",err)
		}
		defer file.Close()
	}
	file,err = os.OpenFile(newImagePath,os.O_WRONLY,0644)
	if err != nil{
		log.Fatalln("error while opening the file..ERROR:",err)
	}
	defer file.Close()
	numwrite, err := file.Write(newImageBytes)
	if err != nil {
		log.Fatalln("error while writing bytes to file..ERROR:",err)
	}
	log.Printf("%v bytes written to %v",numwrite,newImagePath)
	log.Println("reading secrete message...")

}

func createBytes(msg,img []byte, len_msg, num int) []byte{
	arr := []byte{}
	offset := 1
	for i:=1; i <= len_msg; i++ {
		n := num / i
		for j:= offset - 1;j<n;j++ {
			arr = append(arr, img[j])
		}
		arr = append(arr,msg[i-1])
		offset = offset + n
	}
	return arr
}
