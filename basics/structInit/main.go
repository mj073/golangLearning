package main

import (
	"sync"
	"bytes"
	"fmt"
)

type Command struct {
	Init func()
	init sync.Once

	i Info
}

type Info struct {
	v         Vnet
	eventPool sync.Pool
	poller    ifStatsPoller
	pub       *Publisher
}
type Publisher struct {
	Accumulator
	buf *bytes.Buffer
}
type Accumulator struct {
	n              int64
	err            error
	ReaderOrWriter interface{}
}
type ifStatsPoller struct {
	Event
	i            *Info
	sequence     uint
	hwInterfaces ifStatsPollerInterfaceVec
	swInterfaces ifStatsPollerInterfaceVec
	pollInterval float64 // pollInterval in seconds
}
type ifStatsPollerInterfaceVec []ifStatsPollerInterface
type ifStatsPollerInterface struct {
	lastValues map[string]uint64
}
type Event interface {}
type Vnet struct {
	loop Loop
	BufferMain
	cliMain cliMain
	eventMain
	interfaceMain
	packageMain
}
type Loop struct {

}
type BufferMain struct {

}
type cliMain float64
type eventMain string
type interfaceMain int
type packageMain float64
func main(){
	vnet := &Command{
		Init: vnetdInit,
	}
	vnet.i.init()
	fmt.Println(vnet)
	fmt.Println(&vnet.i == vnet.i.poller.i)
}
func (i *Info)init(){
	i.poller.i = i
}
func vnetdInit(){
	fmt.Println("vnetdInit()")
}
