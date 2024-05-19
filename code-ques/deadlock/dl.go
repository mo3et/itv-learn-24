package main

import (
	"fmt"
	"sync"
	"time"
)

/* 这个例子满足了上面的四个条件

对于非缓冲的 go channel，必须是一个goroutine 读，一个 gorouitne 写的，两者存在竞争。这里只有一个主协程，但在 channel 看来 是有两个的一个读，一个写。
读的协程不会放弃 go channel, 会一直占用
读的协程占用的 go channel 没有被别的协程剥夺，(废话这里就一个主协程，拿什么去剥夺)
循环等待，主协程在等待 go channel 里面有东西可以读，而 go channel 等待别人来写，但是就一个主协程，没有人来写。

*/

func mainbad() {
	c1 := make(chan int)
	fmt.Println(<-c1)
}

// Solve
func main() {
	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println(<-c1)
	}()

	c1 <- 1
	wg.Wait()

	goReturn()
}

func goReturn() {
	c := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 1
	}()
	wg.Wait()
	fmt.Println("cgo :", <-c)
}
