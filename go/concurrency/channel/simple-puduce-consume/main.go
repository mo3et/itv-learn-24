package main

import (
	"fmt"
	"time"
)

// simple Producer consumer model

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i // send to channel
		fmt.Println("producer:", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int, done chan bool) {
	for v := range ch {
		fmt.Println("consumer:", v)
		time.Sleep(2 * time.Second)
	}

	done <- true
}

func main() {
	ch := make(chan int, 5) // create buffer channel
	done := make(chan bool)
	go producer(ch)       // start producer goroutine
	go consumer(ch, done) // start consumer goroutine

	// Wait receive from `done`.
	<-done
	fmt.Println("Finish!")
}
