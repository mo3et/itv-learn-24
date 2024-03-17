package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// start goroutine

	go say("world")
	go say("let's Go!")
	time.Sleep(time.Second * 6)
}
