package main

import (
	"github.com/gocql/gocql"
	"log"
)

func main(){
	cluster := gocql.NewCluster("192.168.0.108")
	cluster.Keyspace = "my_cluster"
	session,err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to create cluster session..ERROR:",err)
	}
	defer session.Close()

	log.Println("cluster created and connected successfully")
}
