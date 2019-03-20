package main

import (
	"log"
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

var random = make(chan int)
var tempChan = make(chan int)
type IpFlow struct {
	Id		bson.ObjectId	`bson:"_id"`
	Location_Id	string
	Event_Desc 	string
	Flow_Id		int
	User_Ip 	string
	User_Port 	string
	Destination_Ip	string
	Destination_Port	string
	Ipflow_Timestamp	string
	Record_Count	int
}

func main(){
	ipdr := &IPDR{&DataStore{nil}}
	ipdr.InitDatastore()

	//go ipdr.routine1()
	//go ipdr.routine2()
	//go ipdr.routine3()
	//go ipdr.routine4()
	go ipdr.routine5()
	<- random
}
//func (ipdr *IPDR) routine1(){
//	lp := "routine1: "
//	for {
//		var notifyChan = make(chan int)
//		select {
//		case _ = <-tempChan:
//			_ = ipdr.Ds.IpFlowColl(&notifyChan)
//			log.Println(lp,"never going to be called..")
//		}
//		notifyChan <- 1
//	}
//}
//
//func (ipdr *IPDR) routine2(){
//	lp := "routine2: "
//	for {
//		select {
//		case _ = <-tempChan:
//			_ = ipdr.Ds.FirewallColl()
//			log.Println(lp,"never going to be called..")
//		}
//	}
//}
//func (ipdr *IPDR) routine3(){
//	lp := "routine3: "
//	for {
//		select {
//		case _ = <-tempChan:
//			_ = ipdr.Ds.UUIDColl()
//			log.Println(lp,"never going to be called..")
//		}
//	}
//}
//func (ipdr *IPDR) routine4(){
//	lp := "routine4: "
//	for {
//		select {
//		case _ = <-tempChan:
//			_ = ipdr.Ds.CacheColl()
//			log.Println(lp,"never going to be called..")
//		}
//	}
//}
func (ipdr *IPDR) routine5(){
	lp := "routine1: "
	for {
		ipFlowArr := []IpFlow{}
		userPort := "33345"
		query1 := bson.M{"$and": []interface{}{
			bson.M{"processed":false},
			bson.M{"user_port":userPort},
		}}
		query2 := bson.M{"$and": []interface{}{
			bson.M{"processed":true},
			bson.M{"user_port":userPort},
		}}
		change1 := mgo.Change{
			Update:    bson.M{"$set":bson.M{"processed": true}},
			Upsert: false,
		}
		change2 := mgo.Change{
			Update:    bson.M{"$set":bson.M{"processed": false}},
			Upsert: false,
		}
		log.Println(lp,"fetching ipflow collection...")
		func (){
			//var notifyChan = make(chan int)
			//_ = ipdr.Ds.IpFlowColl(&notifyChan).Find(bson.M{"processed":false}).Sort("ipflow_timestamp").Limit(2000).All(&ipFlowArr)
			//defer func(){notifyChan <- 1}()
			//_ = ipdr.Ds.IpFlowColl().Find(bson.M{"processed":false}).Sort("ipflow_timestamp").Limit(2000).All(&ipFlowArr)

			_ = ipdr.Ds.FindAll(NBRCollIpFlow,query1,[]string{"ipflow_timestamp"},2000,&ipFlowArr)
			log.Println("IpFlow entry:",ipFlowArr)
			go func() {
				<- time.After(time.Second * 10)
				info, err := ipdr.Ds.Update(NBRCollIpFlow,query1,change1)
				if err != nil{
					log.Println("error occured ERROR:",err.Error())
				}else{
					log.Println("info updated:",info.Updated)
					log.Println("IpFlow record updated with {processed:true}")
				}
			}()
			go func() {
				<- time.After(time.Second * 40)
				info, err := ipdr.Ds.UpdateAll(NBRCollIpFlow,query2,change2)
				if err != nil{
					log.Println("error occured ERROR:",err.Error())
				}else{
					log.Println("info updated:",info.Updated)
					log.Println("IpFlow record updated with {processed:false}")
				}
			}()
		}()
		<-time.After(time.Second * 10)
	}
}