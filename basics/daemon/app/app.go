package app

import (
	"basics/daemon/httpServer"
	"basics/daemon/start"
	"basics/daemon/daemons"
	"basics/daemon/parent"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"errors"
)
type cmd interface {
	Main(...string) error
}
type App struct {
	ByName map[string]cmd
}
var Apps = &App{
	ByName: map[string]cmd{
		"start": &start.Command{},
		"daemons": &daemons.Command{
			Daemons: []string{"http","parent"},
		},
		"http": &httpServer.Command{},
		"parent": &parent.Command{},
	},
}
func (app *App) Main(args ...string) error{
	fmt.Println("in app.Main",args)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM)
	defer func(sig chan os.Signal) {
		sig <- syscall.SIGABRT
	}(sig)
	x := os.Args[1]
	v, found := app.ByName[x]
	if !found {
		return errors.New("cmd "+x+" not found")
	}
	v.Main(args...)
	wait(sig)
	fmt.Println("exiting app.Main()")
	return nil
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
}