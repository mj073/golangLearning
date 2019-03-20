package main

import (
	"fmt"
	avro "github.com/dangkaka/go-kafka-avro"
	"github.com/bsm/sarama-cluster"
	"encoding/json"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "hf-counters"

func main() {
	var data map[string]map[string]uint64
	consumerCallbacks := avro.ConsumerCallbacks{
		OnDataReceived: func(msg avro.Message) {
			fmt.Println(msg)
			v := []byte(msg.Value)
			err := json.Unmarshal(v,&data)
			if err != nil{
				panic(err)
			}
			fmt.Println("data:",data)
		},
		OnError: func(err error) {
			fmt.Println("Consumer error", err)
		},
		OnNotification: func(notification *cluster.Notification) {
			fmt.Println(notification)
		},
	}

	consumer, err := avro.NewAvroConsumer(kafkaServers, schemaRegistryServers, topic, "consumer-group", consumerCallbacks)
	if err != nil {
		fmt.Println(err)
	}
	consumer.Consume()
}
