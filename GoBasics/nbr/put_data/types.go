package main

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type Cache struct{
	Id	bson.ObjectId	`bson:"_id"`
	Ingress_Ip 	string
	Ingress_Port	string
	Egress_Ip	string
	Egress_Port	string
	Destination_Ip	string
	Http_Timestamp	string
	Response_Time_Sec	int
	Response_Time_Microsec	int64
	Cache_code	string
	UUID 	string
}
type UUID struct {
	Id	bson.ObjectId	`bson:"_id"`
	UUID 	string
	User_Ip	string
	User_Port	string
	NAT_Ip	string
	NAT_Port	string
	Destination_Ip	string
	Destination_Port	string
	UUID_Timestamp	string
	Processed	bool
}
type Firewall struct{
	Id	bson.ObjectId	`bson:"_id"`
	Ingress_Ip 	string
	Ingress_Port	string
	Public_Ip	string
	Public_Port	string
	Destination_Ip	string
	Destination_Port	string
	Firewall_Timestamp	string
}
type IPDR struct{
	DB_Addr	string
	FilePath	string
	Ds *DataStore
	CacheProcessed chan map[UUID][]Cache
	UUIDProcessed chan UUID
	FirewallProcessed chan Firewall
}

type DataStore struct{
	Session	*mgo.Session
	CacheColl *mgo.Collection
	UUIDColl *mgo.Collection
	FirewallColl *mgo.Collection
}

