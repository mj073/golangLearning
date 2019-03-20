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

	err = ch.ExchangeDeclare("hello-exchange2","direct",true,true,false,false,nil)
	if err != nil{
		log.Fatalln("failed to create exchange..ERROR:",err)
	}

	q, err := ch.QueueDeclare("hello", true, false, false, false, nil)
	if err != nil{
		log.Fatalln("failed to create queue..ERROR:",err)
	}
	err = ch.QueueBind(q.Name,"","hello-exchange2",false,nil)
	if err != nil{
		log.Fatalln("failed to bind a queue to exchange..ERROR:",err)
	}

	q2, err := ch.QueueDeclare("hello2", true, false, false, false, nil)
	if err != nil{
		log.Fatalln("failed to create queue2..ERROR:",err)
	}

	err = ch.QueueBind(q2.Name,"","hello-exchange2",false,nil)
	if err != nil{
		log.Fatalln("failed to bind a queue2 to exchange..ERROR:",err)
	}
	body := "hello-world"
	err = ch.Publish("hello-exchange2","",false,false,amqp.Publishing{ContentType:"text/plain",Body:[]byte(body)})
	if err != nil{
		log.Fatalln("failed to publish a message..ERROR:",err)
	}
}
