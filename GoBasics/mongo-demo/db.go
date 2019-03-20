package main

import (
	"gopkg.in/mgo.v2"
	"log"
)
var(
	NBRDBName = "nbr"
	NBRCollCache = "cache"
	NBRCollUUID = "uuid"
	NBRCollFirewall = "firewall"
	NBRCollIpFlow = "ipflow"
)
func (ipdr *IPDR)InitDatastore() error{
	lp := "InitDataStore: "
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println(lp,"ERROR: failed to connect to db: ", err)
		return err
	}
	log.Println(lp,"connected to db")
	log.Println(lp,"setting db to safe mode")
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
	ipflow_index_meta := mgo.Index{
		Key:        []string{"ipflow_timestamp"},
		Unique:     false,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}

	err = session.Copy().DB(NBRDBName).C(NBRCollCache).EnsureIndex(cache_index_meta)
	if err != nil {
		log.Println(lp,"ERROR: failed to create cache index ERROR:", err)
		return err
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollUUID).EnsureIndex(uuid_index_meta)
	if err != nil {
		log.Println(lp,"ERROR: failed to create uuid index ERROR:", err)
		return err
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollFirewall).EnsureIndex(firewall_index_meta)
	if err != nil {
		log.Println(lp,"ERROR: failed to create firewall index ERROR:", err)
		return err
	}
	err = session.Copy().DB(NBRDBName).C(NBRCollIpFlow).EnsureIndex(ipflow_index_meta)
	if err != nil {
		log.Println(lp,"ERROR: failed to create ipflow index ERROR:", err)
		return err
	}
	ipdr.Ds = &DataStore{session}
	return nil
}

func (mds *DataStore) IpFlowColl(notifyChan *chan int) *mgo.Collection{
//func (mds *DataStore) IpFlowColl() *mgo.Collection{
	s := mds.Session.Copy()
	go func(){
		for {
			log.Println("in for loop")
			select {
			case _ = <- *notifyChan:
				log.Println("session is getting closed now...")
				s.Close()
				goto label
			}
		}
		label:
		log.Println("out of for loop...")
	}()
	return s.DB(NBRDBName).C(NBRCollIpFlow)
}

//func (mds *DataStore) UUIDColl() *mgo.Collection{
//	s := mds.Session.Copy()
//	return s.DB(NBRDBName).C(NBRCollUUID)
//}
//
//func (mds *DataStore) CacheColl() *mgo.Collection{
//	s := mds.Session.Copy()
//	return s.DB(NBRDBName).C(NBRCollCache)
//}
//
//func (mds *DataStore) FirewallColl() *mgo.Collection{
//	s := mds.Session.Copy()
//	return s.DB(NBRDBName).C(NBRCollFirewall)
//}



func (mds *DataStore) FindOne(collName string,query interface{},sortBy []string,limit int,result interface{})(err error){
	s := mds.Session.Copy()
	defer s.Close()
	err = s.DB(NBRDBName).C(collName).Find(query).Sort(sortBy...).Limit(limit).One(result)
	return
}
func (mds *DataStore) FindAll(collName string,query interface{},sortBy []string,limit int,result interface{})(err error){
	s := mds.Session.Copy()
	defer s.Close()
	err = s.DB(NBRDBName).C(collName).Find(query).Sort(sortBy...).Limit(limit).All(result)
	return
}
func (mds *DataStore) Update(collName string,query interface{},change interface{})(info *mgo.ChangeInfo,err error){
	s := mds.Session.Copy()
	defer s.Close()
	info, err = s.DB(NBRDBName).C(collName).Find(query).Apply(change.(mgo.Change),nil)
	return
}
func (mds *DataStore) UpdateAll(collName string,query interface{},change interface{})(info *mgo.ChangeInfo,err error){
	s := mds.Session.Copy()
	defer s.Close()
	info, err = s.DB(NBRDBName).C(collName).UpdateAll(query,change)
	return
}