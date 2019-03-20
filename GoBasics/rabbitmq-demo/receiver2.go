package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main(){
	conn,err := amqp.Dial("amqp://guest:guest@192.168.0.108:5672/")
	if err != nil{
		log.Fatalln("failed to connect to rabbitmq server...ERROR:",err)
	}
	defer conn.Close()

	ch,err := conn.Channel()
	if err != nil{
		log.Fatalln("failed to create channel..ERROR:",err)
	}
	defer ch.Close()

	//q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	//if err != nil{
	//	log.Fatalln("failed to create queue..ERROR:",err)
	//}

	msgs, err := ch.Consume("hello2","consumer2",true,false,false,false,nil)
	if err != nil{
		log.Fatalln("failed to consume messages..ERROR:",err)
	}

	forever := make(chan bool)

	go func() {
		for i := range msgs{
			log.Printf("received message:%s",i.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<- forever
}
