package main

import (
	"aleflibs/messagequeue"
	"fmt"
)

func main(){

	mq :=  &messagequeue.RabbitMq{}
	mq.Init("192.168.0.108:5672")
	//msg := "message"
	mq.BroadCastPublish([]byte("hello"),"alef")
	fmt.Println("message pushed to broker")
}
