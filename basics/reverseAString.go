package main

import (
	"fmt"
	//"os"
)

func main(){
	str := "ma$$hes#h j,,,ad.hav$"
	//str := "mahe#"
	//str := os.Args[1]
	strByte := []byte(str)
	//var newStr string
	//for _,s := range str{
	//	if isAlpha(s) {
	//		newStr = string(s) + newStr
	//	}else {
	//
	//	}
	//}
	//fmt.Println(newStr)
	l := len(str)
	j := l-1
	for i := 0;i<l/2;i++{
		if !isAlpha(strByte[j]){
			j--
			i--
			continue
		}
		if !isAlpha(strByte[i]){
			continue
		}
		strByte[i],strByte[j] = strByte[j],strByte[i]
		j--
	}
	fmt.Println(string(strByte))
}

func isAlpha(s byte) bool{
	if (s >= 97 && s <= 122) || (s >=65 && s <= 90){
		return true
	}
	return false
}