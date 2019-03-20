package main

/*import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"basics/daemon/httpServer"
	"basics/daemon/parent"
	"basics/daemon/start"
	"basics/daemon/daemons"
)
type cmd interface {
	Main(...string) error
}
var Processes = map[string]cmd{
	"start": &start.Command{},
	"daemons": &daemons.Command{
		Daemons: []string{"http","parent"},
	},
	"http": &httpServer.Command{},
	"parent": &parent.Command{},
}
func main() {
	progName := os.Args[0]
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM)
	defer func(sig chan os.Signal) {
		sig <- syscall.SIGABRT
	}(sig)

	x := os.Args[1]
	switch x {
	case "daemons":
		fmt.Println("daemons case")
		v, _ := Processes[x]
		v.Main(os.Args...)
	case "http":
		fmt.Println("http case")
		v, _ := Processes[x]
		v.Main(os.Args...)
	case "parent":
		fmt.Println("parent case")
		v, _ := Processes[x]
		v.Main(os.Args...)
	case "start":
		fmt.Println("start case")
		Processes["start"].Main(progName,"daemons")
	case "stop":

	}
	wait(sig)
}

func wait(ch chan os.Signal) {
	for sig := range ch {
		if sig == syscall.SIGTERM {
			fmt.Println("SIGTERM received")
			os.Stdout.Sync()
			os.Stderr.Sync()
			os.Stdout.Close()
			os.Stderr.Close()
			os.Exit(0)
		}
		break
	}
}*/

import (
	"os"
	"basics/daemon/app"
	"fmt"
)

func main() {
	if err := app.Apps.Main(os.Args...);err != nil{
		fmt.Errorf("%s","failed to start daemon..ERROR"+err.Error())
		os.Exit(1)
	}
}
