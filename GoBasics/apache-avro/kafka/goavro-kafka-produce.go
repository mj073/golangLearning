package main

import (
	"fmt"
	avro "github.com/dangkaka/go-kafka-avro"
	"time"
	"flag"
	"encoding/json"
)
var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "topic1"

func main(){
	var n int
	avroSchema := `
			{"type":"map","values":{"type":"map","values":"long"}}
	`
	//avroSchema := `{"name":"m","type":"map","values":{"type":"map","values":"long"}}`
	producer, err := avro.NewAvroProducer(kafkaServers, schemaRegistryServers)
	if err != nil {
		fmt.Printf("Could not create avro producer: %s", err)
	}
	flag.IntVar(&n, "n", 1, "number")
	flag.Parse()
	for i := 0; i < n; i++ {
		fmt.Println(i)
		addMsg(producer, avroSchema)
	}
}

func addMsg(producer *avro.AvroProducer, schema string) {
	v := map[string]map[string] int64{
		"xeth1": map[string] int64 {
			"counter1": 123,
			"counter2": 456,
		},
	}
	value, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}
	//value := `{"xeth1":{"counter1":123,"counter2":456}}`
	//value := `{"m":{"x":1}}`
	key := time.Now().String()
	err = producer.Add(topic, schema, []byte(key), value)
	fmt.Println(key)
	if err != nil {
		fmt.Printf("Could not add a msg: %s", err)
	}
}

