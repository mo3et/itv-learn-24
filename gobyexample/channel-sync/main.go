package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送 done 通知已经done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// if removed it. will exit before worker run.
	// 阻塞
	<-done
}
