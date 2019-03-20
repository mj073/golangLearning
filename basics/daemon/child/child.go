package child

import (
	"time"
	"fmt"
)

func Init() {
	run()
}

func run() {
	for {
		fmt.Println("chile daemon")
		<- time.After(time.Millisecond * 600)
	}
}