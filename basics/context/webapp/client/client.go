package main

import (
	"net/http"
	"fmt"
	"context"
)

func main(){
	searchUrl := "http://:8080/search?timeout=10"
	resp, err := http.Get(searchUrl)
	if err != nil {
		b := []byte{}
		resp.Body.Read(b)
		fmt.Println(string(b))
	}
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}