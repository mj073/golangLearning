package main

import (
	"net/http"
	"sync"
	"fmt"
)
//var respChan = make(chan string,1)
func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"https://www.google.com/",
		"http://www.somestupidname.com/",
	}
	respChan := make(chan string,1)
	go func (){
		for {
			select {
			case x := <- respChan :
				fmt.Println(x)
			}
		}
	}()
	//go responseReceiveRoutine()
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil && resp != nil{
				respChan <- fmt.Sprintln(url,":",resp.Status)
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
//func responseReceiveRoutine(){
//	for {
//		select {
//		case x := <- respChan :
//			fmt.Println(x)
//		}
//	}
//}