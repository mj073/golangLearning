package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dir := "D:/Calsoft/Platina/SOC dump/th_mem"
	var mem_addr_map  map[string]struct{ addr,block string }
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("err")
	}
	mem_addr_map = make(map[string]struct{ addr,block string })
	for _, f := range files {
		file, err := ioutil.ReadFile(dir+ "/" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		b := string(file)
		for _, v := range strings.Split(b, "\n") {
			if len(v) > 0{
				d := strings.Split(v, " ")
				var name_index int
				var name_block string
				isMemDump := true
				if isMemDump{
					name_index = len(d) - 1
				}else {
					name_index = 2
					name_block = strings.Split(d[name_index],".")
				}

				v := mem_addr_map[name_block[0]]
				v.addr = d[0]
				v.block = name_block[1]
				mem_addr_map[name_block[0]] = v
			}
			//if len(v) > 0{
			//	d := strings.Split(v, " ")
			//
			//	val := mem_addr_map[d[2]]
			//	val.addr = d[0]
			//	mem_addr_map[d[2]] = val
			//}
		}
	}

	for i, j := range mem_addr_map {
		fmt.Println("key:",i,"addr:",j.addr,"block:",j.block)
		goto lable
	}

}