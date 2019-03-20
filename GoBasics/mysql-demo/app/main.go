package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"GoBasics/mysql-demo/utils"
	"fmt"
)

func main() {
	cacheStuct := utils.CacheStruct{}
	db,err := sql.Open("mysql","root:alef123@tcp(192.168.0.108:3306)/nbr?parseTime=true")
	if err != nil {
		log.Println("error while connecting to db server ERROR:",err)
		return
	}
	rows,err := db.Query("SELECT * FROM cache WHERE ingress_ip='192.168.0.108'")
	if err != nil{
		log.Println("error while querying db ERROR:",err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(&cacheStuct.Id,&cacheStuct.IngressIp,&cacheStuct.IngressPort,&cacheStuct.EgressIp,&cacheStuct.EgressPort,&cacheStuct.HttpTimestamp)

		//err := rows.Scan(&cacheStruct)
		if err != nil {
			fmt.Println(err)
		}
		log.Println("cacheStruct: ",cacheStuct)
	}

}
