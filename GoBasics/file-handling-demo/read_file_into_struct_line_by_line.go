package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)
type lossless_config struct{
	shared_limit int
	headroom_limit int
	offset int
	floor int
}

func main(){
	//filename := os.Args[1]
	filename := "D:\\conf.txt"
	conf_bytes, err := ioutil.ReadFile(filename)
	if err != nil{
		panic("failed to read the file...ERROR:"+err.Error())
	}
	fmt.Println("conf_bytes:",conf_bytes)
	conf_raw := string(conf_bytes)
	fmt.Println("conf_raw:",conf_raw)
	var lossless_conf []lossless_config
	for i, s := range strings.Split(conf_raw,"\n"){
		fmt.Printf("s:%v %T i:%v\n",s,s,i)
		var conf lossless_config
		v := strings.Split(s," ")
		fmt.Println("v:",v)
		if len(v) > 1 {
			conf.shared_limit, err = strconv.Atoi(v[0])
			conf.headroom_limit, err = strconv.Atoi(v[1])
			conf.offset, err = strconv.Atoi(v[2])
			conf.floor, err = strconv.Atoi(v[3])
		}
		lossless_conf = append(lossless_conf,conf)
	}
	fmt.Println("lossless_conf:",lossless_conf)
}
