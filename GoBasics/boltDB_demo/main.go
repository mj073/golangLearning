package main

import (
	"github.com/boltdb/bolt"
	"path"
	"log"
	"encoding/json"
	"GoBasics/boltDB_demo/metadata"
)


type BoltDataStore struct{
	db	*bolt.DB
	custBkt string
}
var DS BoltDataStore

func main(){
	log.Println("Initializing BoltDataStore..")
	initDB("D:/boltDB","studentinfo")
	defer closeDB()

	meta := metadata.Metadata{"Mahesh","1234",25}
	meta1 := metadata.Metadata{"mahesh","1235",26}
	meta2 := metadata.Metadata{"MAHESH","1236",28}

	insert(meta)
	insert(meta1)
	insert(meta2)

	log.Println(getMeta("mahesh"))
}
func initDB(filepath, filename string){

	db,err := bolt.Open(path.Join(filepath,filename),0600,nil)
	if err != nil {
		log.Fatal("err: failed to open database:", err)
	}
	log.Println("database ready:", db.Path())
	DS = BoltDataStore{db,"custBkt"}
}
func closeDB(){
	log.Println("Closing bolt data store")
	DS.db.Close()
}
func insert(m metadata.Metadata){
	DS.db.Update(func (tx *bolt.Tx) error{
		b,err := tx.CreateBucketIfNotExists([]byte(DS.custBkt))
		if err != nil {
			log.Println("err: failed to create bucket:", DS.custBkt, err)
			return err
		}
		buf, err := json.Marshal(m)
		if err != nil {
			log.Println("err: failed to marshal metadata", err)
			return err
		}
		return b.Put([]byte(m.Id), buf)
	})
}

func getMeta(identifier string) (meta metadata.Metadata){
	var relm metadata.Metadata
	DS.db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(DS.custBkt))
		c := b.Cursor()

		for k,v := c.First();k!=nil;k,v = c.Next(){

			json.Unmarshal(v,&relm)
			if relm.Name == identifier{
				meta = relm
			}
		}
		return nil
	})
	return
}