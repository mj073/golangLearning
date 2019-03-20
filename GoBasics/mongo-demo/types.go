package main

import "gopkg.in/mgo.v2"

type IPDR struct {
	Ds *DataStore
}

type DataStore struct {
	Session *mgo.Session
}
