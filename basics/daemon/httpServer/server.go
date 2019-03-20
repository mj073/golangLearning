package httpServer

import (
	"net/http"
	"log"
	"basics/daemon/daemons"
	"fmt"
)

type Command struct {
	daemons *daemons.Command
}
func (cmd *Command) Main(s ...string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/",SayHelloWorld)
	go func() {
		log.Fatalln(http.ListenAndServe(":8080", mux))
	}()
	return nil
}
func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world API called")
	html := "Hello World"

	w.Write([]byte(html))
}