package rpcobject

import (
	"sync"
	"fmt"
)

type Agent struct {
	Name string
	KafkaClient *KafkaClient
	RpcServer *RpcServer
}
type KafkaClient struct {
	Agent *Agent
	ServerIP string
	ServerPort string
	Topic string
}
type RpcServer struct {
	Agent *Agent
	Request *RpcRequest
	mu *sync.RWMutex
}
type RpcRequest struct {
	GetAgent int32
}
func newKafkaClient(a interface{}) *KafkaClient{
	k := &KafkaClient{}
	k.ServerIP = "localhost"
	k.ServerPort = "9092"
	k.Topic = "agent"
	k.Agent = a.(*Agent)
	return k
}
func NewAgent() *Agent{
	a := &Agent{}
	a.Name = "rpcAgent"
	a.KafkaClient = newKafkaClient(a)
	a.RpcServer = &RpcServer{Agent: a, Request: &RpcRequest{}, mu: &sync.RWMutex{}}
	return a
}
func (r *RpcServer) GetAgent(args string, reply *Agent) error {
	var a *Agent
	fmt.Println("GetAgent agent:",r.Agent)
	reply = r.Agent
	a = r.Agent
	fmt.Println("GetAgent a:",a)
	return nil
}

