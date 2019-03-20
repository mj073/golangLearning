package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
	"os"
	"flag"
	"log"
)

type Cache struct{
	Ingress_Ip 	string
	Ingress_Port	string
	Egress_Ip	string
	Egress_Port	string
	Destination_Ip	string
	Http_Timestamp	string
	Response_Time_Sec	int
	Response_Time_Microsec	int64
	Cache_Code	string
}
type UUID struct {
	UUID 	string
	User_Ip	string
	User_Port	string
	NAT_Ip	string
	NAT_Port	string
	Destination_Ip	string
	Destination_Port string
	UUID_Timestamp	string
}
type Firewall struct{
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
}

var(
	NBRDBName = "nbr"
	NBRCollCache = "cache"
	NBRCollUUID = "uuid"
	NBRCollFirewall = "firewall"
)
const (
	in_time_format = "2006-01-02 15:04:05"
	out_time_format= "02/01/2006 15:04:05"
)

func main() {
	var ipdr IPDR

	flag.StringVar(&ipdr.DB_Addr, "db", "", "MongoDB IP address")
	flag.StringVar(&ipdr.FilePath, "file", "", "File location to store ipdr records")
	flag.Parse()

	final_record := []string{}
	session, err := mgo.Dial(ipdr.DB_Addr)
	if err != nil {
		log.Fatalln("ERROR: failed to connect to db: ", err)
	}
	log.Println("connected to db")

	log.Println("setting db to safe mode")
	session.SetSafe(&mgo.Safe{})

	// Set unique and other constraints
	cache_index_meta := mgo.Index{
		Key:        []string{"ingress_ip","ingress_port"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}
	uuid_index_meta := mgo.Index{
		Key:        []string{"nat_ip","nat_port"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}
	firewall_index_meta := mgo.Index{
		Key:        []string{"ingress_ip","ingress_port"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollCache).EnsureIndex(cache_index_meta)
	if err != nil {
		log.Fatalln("err: failed to create cache index ERROR:", err)
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollUUID).EnsureIndex(uuid_index_meta)
	if err != nil {
		log.Fatalln("err: failed to create uuid index ERROR:", err)
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollFirewall).EnsureIndex(firewall_index_meta)
	if err != nil {
		log.Fatalln("err: failed to create firewall index ERROR:", err)
	}
	cacheColl := session.DB(NBRDBName).C(NBRCollCache)
	uuidColl := session.DB(NBRDBName).C(NBRCollUUID)
	firewallColl := session.DB(NBRDBName).C(NBRCollFirewall)

	cache := []Cache{}
	err = cacheColl.Find(nil).All(&cache)
	if err != nil{
		log.Fatalln("error while fetching cache data from db ERROR:",err)
		return
	}
	//fmt.Println("current time:",time.Now().Unix()," Time after 10 seconds: ",time.Now().Unix()+10)
	//fmt.Println("current time:",time.Now().UnixNano()," Time after 10 nanoseconds: ",time.Now().UnixNano()+10)
	log.Println("cache length:",len(cache))
	for _,c := range(cache){
		var startTimeStr string
		var endTimeStr string

		uuid := []UUID{}
		firewall := []Firewall{}
		u := UUID{}
		f := Firewall{}
		record := ""

		err = uuidColl.Find(bson.M{"$and": []interface{}{
			bson.M{"nat_ip":c.Ingress_Ip},
			bson.M{"nat_port":c.Ingress_Port},
			//bson.M{"destination_ip":c.Destination_Ip},
			bson.M{"uuid_timestamp":bson.M{"$lte": c.Http_Timestamp}},
		}}).All(&uuid)
		if err != nil {
			log.Println("error while fetching uuid data from db ERROR:",err)
			return
		}
		if len(uuid) > 1 {
			//log.Println("uuid: multiple entrys for NAT-IP NAT-PORT are present")
			//continue
			u = uuid[0]
			startTimeStr = u.UUID_Timestamp
			//record = record + fmt.Sprint(startTimeStr,",",u.User_Ip,",",u.User_Port,",")
		}else if len(uuid) == 1 {
			u = uuid[0]
			startTimeStr = u.UUID_Timestamp
			//record = record + fmt.Sprint(startTimeStr,",",u.User_Ip,",",u.User_Port,",")
		}else {
			log.Printf("uuid logs for NAT_IP:NAT_PORT %s:%s has not generated yet..http_timestamp is %s..destination_ip is %s\n",c.Ingress_Ip,c.Ingress_Port,c.Http_Timestamp,c.Destination_Ip)
		}

		startTime, err := time.Parse(in_time_format,startTimeStr)
		if err != nil {
			log.Fatalln("error while parsing start date time ERROR:",err)
			return
		}
		startTimeStr = startTime.Format(out_time_format)
		record = record + fmt.Sprint(startTimeStr,",",u.User_Ip,",",u.User_Port,",")

		responseTimeSec, err := time.ParseDuration(strconv.Itoa(c.Response_Time_Sec)+"s")
		if err != nil {
			log.Fatalln("error while parsing responseTimeSec duration ERROR:",err, responseTimeSec)
			return
		}
		//responseTimeMicrosec, err := time.ParseDuration(c.Response_Time_Microsec)
		//if err != nil {
		//	log.Fatalln("error while parsing responseTimeMicrosec duration ERROR:",err)
		//	return
		//}
		endTimeStr = startTime.Add(responseTimeSec).Format(out_time_format)

		if c.Cache_Code != "0x0001"{
			err = firewallColl.Find(bson.M{"$and": []interface{}{
				bson.M{"ingress_ip":c.Egress_Ip},
				bson.M{"ingress_port":c.Egress_Port},
				bson.M{"destination_ip":c.Destination_Ip},
			}}).All(&firewall)

			if len(firewall) > 1 {
				//log.Println("firewall: logic for correlating multiple nat_ip nat_port combination yet to be implemented")
				f = firewall[0]
				record = record + fmt.Sprint(f.Public_Ip,",",f.Public_Port,",",f.Destination_Ip,",",f.Destination_Port,",")
			}else if len(firewall) == 1 {
				f = firewall[0]
				record = record + fmt.Sprint(f.Public_Ip,",",f.Public_Port,",",f.Destination_Ip,",",f.Destination_Port,",")
			}else {
				//log.Printf("firewall logs for Egress_Ip:Egress_Port %s:%s has not generated yet\n",c.Egress_Ip,c.Egress_Port)
				continue
			}
			record = record + fmt.Sprint(endTimeStr)
		}else{
			//log.Println("the request has been served from cache")
			record = record + fmt.Sprint("0.0.0.0",",","0",",",u.Destination_Ip,",",u.Destination_Port,",",endTimeStr)
		}
		final_record = append(final_record,record)
	}
	err = WriteToFile(&ipdr,final_record)
	if err != nil {
		log.Fatalln("Error while writing to file")
	}
}

func WriteToFile(ipdr *IPDR,final_record []string) (err error){
	if _, err := os.Stat(ipdr.FilePath); os.IsNotExist(err) {
		log.Println("file:",ipdr.FilePath," does not exist...creating file")
		_, err = os.Create(ipdr.FilePath)
		if err != nil {
			return err
		}
		log.Println("file:",ipdr.FilePath," created successfully")
	}
	log.Println("opening file:",ipdr.FilePath," in append mode")
	f, err := os.OpenFile(ipdr.FilePath, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.Println("file opened in append mode")
	defer f.Close()
	log.Println("appending record to the file")
	for i,log := range(final_record){
		//fmt.Println(log)
		log = strconv.Itoa(i+1) + "," + log + "\n"
		if _, err = f.WriteString(log); err != nil {
			return err
		}
	}
	return nil
}