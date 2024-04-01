package main

import (
	"context"
	"fmt"
	"time"
)



// 实现 Cancel Signal

func ContextCancelExample() {
	message := make(chan int, 10)

	// producer
	for i := 0; i < 10; i++ {
		message <- i
	}

	// Timeout时自动执行cancel，ctx会发送ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Define a new type for the key
	type UserIDKey int

	// Use the new type as the key
	userid := UserIDKey(10)
	user := "jason"
	ctx = context.WithValue(ctx, userid, user)
	// consumer
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)

		// 要实现闭包效果 需要将闭包函数赋予一个变量，相当于初始化
		// 每次调用就会叠加
		val := ctx.Value(userid)
		helloFunc := Hello(val.(string))

		for range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-message)
				// 之前错误原因在于多次实例化 Hello(val),导致被覆盖
				helloFunc()
			}
		}
	}(ctx)

	defer close(message)
	// 防御式编程 如果在超时前完成则释放资源
	defer cancel()

	<-ctx.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}

func main() {
	ContextCancelExample()
}

func Hello(nameStr string) func() {
	// nameStr := name.(string)
	s := 0
	if nameStr == "jason" {
		return func() {
			s++
			fmt.Printf("%d is %s\n", s, nameStr)
		}
	}
	return nil
}
