package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	for {
		val, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(val)
		ch2 <- val
	}
	go func() {
		for {
			fmt.Println(<-ch2)
		}
	}()
}
