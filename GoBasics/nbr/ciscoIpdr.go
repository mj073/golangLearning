package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"time"
	l "github.com/Sirupsen/logrus"
	"github.com/magiconair/properties"
)
var (
	NBRDBName = "ciscoIpdr"
	NBRCollIpFlow = "ipflow"
	random chan bool
	LOG_LEVEL = l.InfoLevel
	ipdrlog *l.Entry
)
type IpFlow struct {
	Id		bson.ObjectId	`bson:"_id"`
	Event_Desc 	string
	Flow_Id		int
	User_Ip 	string
	User_Port 	string
	Destination_Ip	string
	Destination_Port	string
	Ipflow_Timestamp	string
}
type Ipdr struct{
	IpFlowProcessed chan IpFlow
	Ds *DataStore
}
type DataStore struct{
	IpFlowColl *mgo.Collection
}
func init(){
	ipdrlog = l.WithField("ipdr","cisco")
	ipdrlog.Level = LOG_LEVEL
}
func main(){
	ipdr := Ipdr{make(chan IpFlow),nil}
	ipdr.InitDataStore("192.168.0.108")
	go ipdr.checkLogLevel()
	go func(){
		for {
			ipdrlog.Println("in infinite go routine...")
			time.Sleep(time.Second * 20)
			ipdrlog.Debugln("in debug of inifinite go routine...")
		}
	}()
	//go ipdr.EndOrInvalidateEvent()
	//go ipdr.StartIpdrRoutine()
	//go ipdr.UpdateIpFlowRecordRoutine()
	<- random
}
func (ipdr *Ipdr)checkLogLevel(){
	lp := "checkLogLevel: "
	ticker := time.NewTicker(time.Second * 10)
	var new_log_level string
	for range ticker.C{
		prop := properties.MustLoadFile("E:/golang/src/GoBasics/nbr/resources.properties",properties.UTF8)
		new_log_level = prop.MustGetString("LOG_LEVEL")
		if new_log_level != LOG_LEVEL.String(){
			ipdrlog.Println(lp,"log level changed to ",new_log_level)
			switch new_log_level{
			case "info":
				LOG_LEVEL = l.InfoLevel
			case "debug":
				LOG_LEVEL = l.DebugLevel
			case "warning":
				LOG_LEVEL = l.WarnLevel
			case "error":
				LOG_LEVEL = l.ErrorLevel
			case "fatal":
				LOG_LEVEL = l.FatalLevel
			case "panic":
				LOG_LEVEL = l.PanicLevel
			}
			ipdrlog.Level = LOG_LEVEL
		}
	}
}
func (ipdr *Ipdr) EndOrInvalidateEvent(){
	lp := "EndOrInvalidateEvent: "
	ipFlowColl := ipdr.Ds.IpFlowColl
	ipFlow := []IpFlow{}
	err := ipFlowColl.Find(bson.M{"$and": []interface{}{
		bson.M{"$or":[]interface{}{
			bson.M{"event_desc":"FLOW_END_EVENT"},
			bson.M{"event_desc":"FLOW_INVALIDATE_EVENT"},
		}},
		bson.M{"processed":false},
	}}).Sort("ipflow_timestamp").Limit(200).All(&ipFlow)
	if err != nil{
		ipdrlog.Errorln(lp,"ERROR: error while fetching flow data from db ERROR:",err)
	}
	ipdrlog.Println(lp,"len of EndOrInvalidateEvent:",len(ipFlow))
	//for _,flow := range(ipFlow){
	//
	//}
}
func (ipdr *Ipdr) StartIpdrRoutine(){
	lp := "StartIpdrRoutine: "
	ipdrlog.Println(lp,"running the routine every 10 seconds..")
	ticker := time.NewTicker(time.Second * 10)
	for range ticker.C{
		execute(ipdr)
	}
}
func execute(ipdr *Ipdr){
	lp := "StartIpdrRoutine_execute: "
	ipflows := []IpFlow{}
	ipdrlog.Debugln(lp,"fetching 200 ipflow records from db...")
	err := ipdr.Ds.IpFlowColl.Find(bson.M{"processed":false}).Sort("ipflow_timestamp").Limit(200).All(&ipflows)
	if err != nil{
		ipdrlog.Errorln(lp,"ERROR: error while fetching flow data from db ERROR:",err)
		return
	}
	ipdrlog.Println(lp,"ip flow data length:",len(ipflows))
	ipdrlog.Debugln(lp,"iteraing over all the ipflows data...")
	for _,flow := range(ipflows){
		record := ""
		switch flow.Event_Desc {
		case "FLOW_START_EVENT":
			ipdrlog.Println(lp,"start event occurred..[",flow,"]")
			record = flow.Ipflow_Timestamp + "," + flow.User_Ip + "," + flow.User_Port + "," + flow.Destination_Ip + "," + flow.Destination_Port + "," + "271"
			ipdrlog.Println(lp,"record generated: [",record,"]")
			ipdr.IpFlowProcessed <- flow
		case "FLOW_END_EVENT":
			ipdrlog.Println(lp,"end event occurred..[",flow,"]")
			record = flow.Ipflow_Timestamp + "," + flow.User_Ip + "," + flow.User_Port + "," + flow.Destination_Ip + "," + flow.Destination_Port + "," + "272"
			ipdrlog.Println(lp,"record generated: [",record,"]")
			ipdr.IpFlowProcessed <- flow
		case "FLOW_INVALIDATE_EVENT":
			open_flows := []IpFlow{}
			err := ipdr.Ds.IpFlowColl.Find(bson.M{"$and": []interface{}{
				bson.M{"event_desc": "FLOW_START_EVENT"},
				bson.M{"processed": true},
				bson.M{"ended": false},
				bson.M{"ipflow_timestamp": bson.M{"$lte": flow.Ipflow_Timestamp}},
			}}).Sort("ipflow_timestamp").All(&open_flows)
			ipdrlog.Println(lp,"length of open_flows:",len(open_flows))
			if err != nil{
				ipdrlog.Errorln(lp,"ERROR: error while fetching flow data from db ERROR:",err)
				ipdrlog.Debugln(lp,"continuing with next ipflow data...")
				continue
			}
			for _,f := range (open_flows){
				record = f.Ipflow_Timestamp + "," + f.User_Ip + "," + f.User_Port + "," + f.Destination_Ip + "," + f.Destination_Port + "," + "272"
				ipdrlog.Println(lp,"record generated: [",record,"]")
				//ipdr.IpFlowProcessed <- flow
			}
			ipdr.IpFlowProcessed <- flow
		default:
		}
	}
}
func (ipdr *Ipdr)InitDataStore(ip string){
	lp := "InitDataStore: "
	ipdrlog.Debugln(lp,"dialing in to mongo server...")
	session, err := mgo.Dial(ip)
	if err != nil {
		log.Fatalln(lp,"ERROR: failed to connect to db: ", err)
	}
	ipdrlog.Println(lp,"connected to db")
	ipdrlog.Println(lp,"setting db to safe mode")
	session.SetSafe(&mgo.Safe{})

	// Set unique and other constraints
	//cache_index_meta := mgo.Index{
	//	Key:        []string{"ingress_ip","ingress_port"},
	//	Unique:     false,
	//	DropDups:   false,
	//	Background: true, // See notes.
	//	Sparse:     true,
	//}
	//uuid_index_meta := mgo.Index{
	//	Key:        []string{"nat_ip","nat_port"},
	//	Unique:     false,
	//	DropDups:   false,
	//	Background: true, // See notes.
	//	Sparse:     true,
	//}
	//firewall_index_meta := mgo.Index{
	//	Key:        []string{"ingress_ip","ingress_port"},
	//	Unique:     false,
	//	DropDups:   false,
	//	Background: true, // See notes.
	//	Sparse:     true,
	//}
	ipflow_index_meta := mgo.Index{
		Key:        []string{"ipflow_timestamp"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}

	//err = session.Copy().DB(NBRDBName).C().EnsureIndex(cache_index_meta)
	//if err != nil {
	//	log.Fatalln(lp,"ERROR: failed to create cache index ERROR:", err)
	//}
	//err = session.Copy().DB(NBRDBName).C(NBRCollUUID).EnsureIndex(uuid_index_meta)
	//if err != nil {
	//	log.Errorln(lp,"ERROR: failed to create uuid index ERROR:", err)
	//	return err
	//}
	//err = session.Copy().DB(NBRDBName).C(NBRCollFirewall).EnsureIndex(firewall_index_meta)
	//if err != nil {
	//	log.Errorln(lp,"ERROR: failed to create firewall index ERROR:", err)
	//	return err
	//}
	err = session.Copy().DB(NBRDBName).C(NBRCollIpFlow).EnsureIndex(ipflow_index_meta)
	if err != nil {
		ipdrlog.Fatalln(lp,"ERROR: failed to create ipflow index ERROR:", err)
	}

	//ipdr.Ds = &DataStore{session,session.DB(NBRDBName).C(NBRCollCache),session.DB(NBRDBName).C(NBRCollUUID),session.DB(NBRDBName).C(NBRCollFirewall),session.DB(NBRDBName).C(NBRCollIpFlow)}
	ipdr.Ds = &DataStore{session.DB(NBRDBName).C(NBRCollIpFlow)}
}
func (ipdr *Ipdr) UpdateIpFlowRecordRoutine(){
	lp := "UpdateIpFlowRecordRoutine: "
	for {
		if ipFlowProcessed, ok := <- ipdr.IpFlowProcessed; ok{
			ipdrlog.Debugln(lp,"updating ipflow with Flow_ID:",ipFlowProcessed.Flow_Id,"as processed...")
			change := mgo.Change{
				Update:    bson.M{"$set":bson.M{"processed": true}},
				Upsert:    false,
				ReturnNew: true,
			}
			change2 := mgo.Change{
				Update:    bson.M{"$set":bson.M{"ended": true,"time": time.Now()}},
				Upsert:    false,
				ReturnNew: true,
			}
			change3 := bson.M{"$set":bson.M{"ended": true,"time": time.Now()}}
			if ipFlowProcessed.Id != "" {
				_, err := ipdr.Ds.IpFlowColl.Find(bson.M{"_id":ipFlowProcessed.Id}).Apply(change, nil)
				if err != nil {
					ipdrlog.Errorln(lp,"failed updating processed IpFlow record with ObjectId:", ipFlowProcessed.Id, " ERROR:", err)
					//failedId = f.Id.String()
				}
			}
			if ipFlowProcessed.Event_Desc == "FLOW_END_EVENT" {
				ipdrlog.Println(lp,"updating end status for FLOW_START_EVENT...")
				_, err := ipdr.Ds.IpFlowColl.Find(bson.M{"$and": []interface{}{
						bson.M{"event_desc": "FLOW_START_EVENT"},
						bson.M{"ended": false},
						bson.M{"user_ip":ipFlowProcessed.User_Ip},
						bson.M{"user_port":ipFlowProcessed.User_Port},
						bson.M{"destination_ip":ipFlowProcessed.Destination_Ip},
						bson.M{"destination_port":ipFlowProcessed.Destination_Port},
						bson.M{"ipflow_timestamp": bson.M{"$lte": ipFlowProcessed.Ipflow_Timestamp}},
					}}).Apply(change2, nil)
				if err != nil {
					ipdrlog.Errorln(lp,"failed updating processed IpFlow record with ObjectId:", ipFlowProcessed.Id, " ERROR:", err)
				}
			}else if ipFlowProcessed.Event_Desc == "FLOW_INVALIDATE_EVENT"{
				ipdrlog.Println(lp,"FLOW_INVALIDATE_EVENT occured...")
				selector := bson.M{"$and": []interface{}{
						bson.M{"event_desc": "FLOW_START_EVENT"},
						bson.M{"processed": true},
						bson.M{"ended": false},
						bson.M{"ipflow_timestamp": bson.M{"$lte": ipFlowProcessed.Ipflow_Timestamp}},
					}}
				_,err := ipdr.Ds.IpFlowColl.UpdateAll(selector,change3)
				if err != nil {
					ipdrlog.Errorln(lp,"failed updating end status of IpFlow record with FlowID:",ipFlowProcessed.Flow_Id," ERROR:", err)
					//failedId = f.Id.String()
				}
			}
		}
	}
}