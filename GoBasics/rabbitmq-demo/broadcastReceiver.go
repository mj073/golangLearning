package main

import (
	"aleflibs/messagequeue"
	"fmt"
	"strconv"
	"github.com/streadway/amqp"
	"encoding/base64"
	"encoding/json"
)

var task chan string
var mq = &messagequeue.RabbitMq{}
var loop chan bool
func main() {
	fmt.Println("initializing the message bus...")
	mq.Init("192.168.0.108:5672")
	for i := 0 ; i<2 ; i++ {
		fmt.Println("subscribing for queue_",i)
		mq.SubscribeForBrodcastMessaging("alef","queue_"+strconv.Itoa(i))
		go receiver(i,mq)
	}

}

func receiver(i int,mq messagequeue.Message_queue){
	x := i
	fmt.Println("x: ",x)
	b := []byte{}
	c := []byte{}
	for {
		if t,ok := (<-mq.ReceiveBroadcastMessages()).(<- chan amqp.Delivery);ok {
			for m := range t {
				_,err := base64.StdEncoding.Decode(b, m.Body)
				if err != nil{
					fmt.Println("wgvqwerfg")
				}
			}
		}
		fmt.Println("ans: ", string(b))
	}
}
