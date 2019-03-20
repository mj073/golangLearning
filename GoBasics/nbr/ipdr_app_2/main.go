package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
	"log"
	"os"
	"flag"
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
	Session	*mgo.Session
}
var(
	NBRDBName = "nbr"
	NBRCollCache = "cache"
	NBRCollUUID = "uuid"
	NBRCollFirewall = "firewall"
	err error
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
	ipdr.Session, err = mgo.Dial(ipdr.DB_Addr)
	if err != nil {
		log.Fatalln("ERROR: failed to connect to db: ", err)
	}
	log.Println("connected to db")
	session := ipdr.Session
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
		log.Println("ERROR: failed to create cache index ERROR:", err)
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollUUID).EnsureIndex(uuid_index_meta)
	if err != nil {
		log.Println("ERROR: failed to create uuid index ERROR:", err)
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollFirewall).EnsureIndex(firewall_index_meta)
	if err != nil {
		log.Println("ERROR: failed to create firewall index ERROR:", err)
	}
	uuid := []UUID{}
	uuid_repeated := []UUID{}
	uuid_processed := []UUID{}
	firewall_processed := []Firewall{}

	cacheColl := session.DB(NBRDBName).C(NBRCollCache)
	uuidColl := session.DB(NBRDBName).C(NBRCollUUID)
	firewallColl := session.DB(NBRDBName).C(NBRCollFirewall)

	err = uuidColl.Find(bson.M{"processed":false}).Sort("uuid_timestamp").Limit(2000).All(&uuid)
	if err != nil{
		log.Println("ERROR: error while fetching uuid data from db ERROR:",err)
		return
	}

	log.Println("uuid length:",len(uuid))
	for _,u := range(uuid){
		err = uuidColl.Find(bson.M{"$and": []interface{}{
			bson.M{"nat_ip":u.NAT_Ip},
			bson.M{"nat_port": u.NAT_Port},
			bson.M{"processed": false},
		}}).Sort("uuid_timestamp").All(&uuid_repeated)
		if err != nil{
			log.Println("ERROR: error while fetching repeated uuid data from db ERROR:",err)
			log.Printf("Skipping the uuid record having nat_ip:nat_port %s:%s and uuid_timestamp %s\n",u.NAT_Ip,u.NAT_Port,u.UUID_Timestamp)
			continue
		}

		var startTimeStr string
		var endTimeStr string

		cache := []Cache{}
		cache_processed := []Cache{}

		for i := range(uuid_repeated){
			if uuid_repeated[i].UUID == u.UUID && uuid_repeated[i].UUID_Timestamp == u.UUID_Timestamp {
				if i == len(uuid_repeated) - 1 {
					err = cacheColl.Find(bson.M{"$and": []interface{}{
						bson.M{"uuid": ""},
						bson.M{"ingress_ip": u.NAT_Ip},
						bson.M{"ingress_port": u.NAT_Port},
						bson.M{"http_timestamp":bson.M{"$gte": uuid_repeated[i].UUID_Timestamp}},
					}}).Sort("http_timestamp").All(&cache)
				} else {
					err = cacheColl.Find(bson.M{"$and": []interface{}{
						bson.M{"uuid": ""},
						bson.M{"ingress_ip": u.NAT_Ip},
						bson.M{"ingress_port": u.NAT_Port},
						bson.M{"http_timestamp":bson.M{
							"$gte": uuid_repeated[i].UUID_Timestamp,
							"$lt": uuid_repeated[i + 1].UUID_Timestamp,
						}},
					}}).Sort("http_timestamp").All(&cache)
				}
				startTimeStr = u.UUID_Timestamp
			}
		}
		if err != nil {
			log.Println("ERROR: error while fetching cache data from db ERROR:",err)
			log.Printf("Skipping the uuid record having nat_ip:nat_port %s:%s and uuid_timestamp %s\n",u.NAT_Ip,u.NAT_Port,u.UUID_Timestamp)
			continue
		}
		for _,c := range(cache) {
			record := ""
			f := Firewall{}

			startTime, err := time.Parse(in_time_format,startTimeStr)
			if err != nil {
				log.Println("ERROR: error while parsing start date time ERROR:",err)
				log.Printf("Skipping the cache record having ingress_ip:ingress_port %s:%s and http_timestamp %s\n",c.Ingress_Ip,c.Ingress_Port,c.Http_Timestamp)
				continue
			}
			startTimeString := startTime.Format(out_time_format)

			record = record + fmt.Sprint(startTimeString,",",u.User_Ip,",",u.User_Port,",")

			if c.Cache_code != "0x0001" {
				err = firewallColl.Find(bson.M{"$and": []interface{}{
					bson.M{"ingress_ip":c.Egress_Ip},
					bson.M{"ingress_port":c.Egress_Port},
					bson.M{"destination_ip":u.Destination_Ip},
					bson.M{"destination_port":u.Destination_Port},
					bson.M{"firewall_timestamp": bson.M{"$gte":c.Http_Timestamp}},
				}}).Sort("firewall_timestamp").One(&f)

				//if firewall log not found in db
				if f.Firewall_Timestamp != "" {
					record = record + fmt.Sprint(f.Public_Ip, ",", f.Public_Port, ",", f.Destination_Ip, ",", f.Destination_Port, ",")
					endTime, err := time.Parse(in_time_format, f.Firewall_Timestamp)
					if err != nil {
						log.Println("ERROR: error while parsing end date time ERROR:", err)
						log.Printf("Skipping the cache record having ingress_ip:ingress_port %s:%s and http_timestamp %s\n",c.Ingress_Ip,c.Ingress_Port,c.Http_Timestamp)
						continue
					}
					endTimeStr := endTime.Format(out_time_format)
					record = record + fmt.Sprint(endTimeStr)
				}else {
					log.Println("request not served through cache...firewall log not found..")
					responseTimeSec, err := time.ParseDuration(strconv.Itoa(c.Response_Time_Sec) + "s")
					if err != nil {
						log.Println("ERROR: error while parsing responseTimeSec duration ERROR:", err, responseTimeSec)
						log.Printf("Skipping the cache record having ingress_ip:ingress_port %s:%s and http_timestamp %s\n",c.Ingress_Ip,c.Ingress_Port,c.Http_Timestamp)
						continue
					}
					endTimeStr = startTime.Add(responseTimeSec).Format(out_time_format)
					record = record + fmt.Sprint("0.0.0.0", ",", "0", ",", u.Destination_Ip, ",", u.Destination_Port, ",", endTimeStr)
				}
			} else {
				responseTimeSec, err := time.ParseDuration(strconv.Itoa(c.Response_Time_Sec) + "s")
				if err != nil {
					log.Println("ERROR: error while parsing responseTimeSec duration ERROR:", err, responseTimeSec)
					log.Printf("Skipping the cache record having ingress_ip:ingress_port %s:%s and http_timestamp %s\n",c.Ingress_Ip,c.Ingress_Port,c.Http_Timestamp)
					continue
				}
				endTimeStr = startTime.Add(responseTimeSec).Format(out_time_format)
				//log.Println("the request has been served from cache")
				record = record + fmt.Sprint("0.0.0.0", ",", "0", ",", u.Destination_Ip, ",", u.Destination_Port, ",", endTimeStr)
			}
			final_record = append(final_record,record)
			cache_processed = append(cache_processed,c)
			firewall_processed = append(firewall_processed,f)
		}
		uuid_processed = append(uuid_processed,u)
		var failedIds []string
		log.Println("updating ",len(cache_processed)," cache records")
		for _,c := range(cache_processed) {
			id, err := UpdateProcessedCacheRecord(cacheColl, c, u)
			if err != nil {
				log.Printf("ERROR: error while updating processed cache records...ObjectIds: %v\n", id)
				failedIds = append(failedIds,id)
			}
		}
		log.Println("updated ",len(cache_processed)-len(failedIds)," cache records")
	}
	err = WriteToFile(&ipdr,final_record)
	if err != nil {
		log.Fatalln("ERROR: error while writing to file ERROR:",err)
	}

	var failedIds []string
	log.Println("updating ",len(uuid_processed)," uuid records")
	for _,u := range(uuid_processed) {
		id, err := UpdateProcessedUUIDs(uuidColl, u)
		if err != nil {
			log.Printf("ERROR: error while updating processed uuids...ObjectIds: %v\n", id)
			failedIds = append(failedIds,id)
		}
	}
	log.Println("updated ",len(uuid_processed)-len(failedIds)," uuid records")

	var ids []string
	log.Println("updating ",len(firewall_processed)," firewall records")
	for _,f := range(firewall_processed) {
		id, err := UpdateProcessedFirewallRecord(firewallColl, f)
		if err != nil {
			log.Printf("ERROR: error while updating processed firewall records...ObjectIds: %v\n", id)
			ids = append(ids,id)
		}
	}
	log.Println("updated ",len(firewall_processed)-len(ids)," firewall records")
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
	for _,record := range(final_record){
		record = record + "\n"
		if _, err = f.WriteString(record); err != nil {
			return err
		}
	}
	return nil
}

func UpdateProcessedUUIDs(uuidColl *mgo.Collection, u UUID) (failedId string,err error){
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"processed": true}},
		Upsert:    false,
		ReturnNew: true,
	}
	_, err = uuidColl.Find(bson.M{"_id":u.Id}).Apply(change,nil)
	if err != nil {
		log.Println("ERROR: failed updating processed UUIDs: ", err)
		failedId = u.Id.String()
	}
	return
}

func UpdateProcessedCacheRecord(cacheColl *mgo.Collection, c Cache,u UUID) (failedId string,err error){
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"uuid": u.UUID}},
		Upsert:    false,
		ReturnNew: true,
	}
	_, err = cacheColl.Find(bson.M{"_id":c.Id}).Apply(change,nil)
	if err != nil {
		log.Println("ERROR: failed updating processed cache records: ", err)
		failedId = c.Id.String()
	}
	return
}

func UpdateProcessedFirewallRecord(firewallColl *mgo.Collection, f Firewall) (failedId string,err error){
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"processed": true}},
		Upsert:    false,
		ReturnNew: true,
	}
	_, err = firewallColl.Find(bson.M{"_id":f.Id}).Apply(change,nil)
	if err != nil {
		log.Println("ERROR: failed updating processed firewall records: ", err)
		failedId = f.Id.String()
	}
	return
}