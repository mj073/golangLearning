package parent

import (
	"fmt"
	"time"
	"basics/daemon/daemons"
)
type Command struct {
	daemons *daemons.Command
}
func (c *Command) Main(s ...string) error {
	initParent()
	return nil
}
func initParent() {
	go run()
}

func run() {
	for {
		fmt.Println("parent daemon")
		<- time.After(time.Millisecond * 600)
	}
}