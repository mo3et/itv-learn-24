package main

import (
	"fmt"
	"sync"
	"time"
)

// N:1 Producer-Consumer Model
// https://cloud.tencent.com/developer/article/2301609

const rateLimit = 5

type msgChan struct { // 用于通信的结构体
	Id       int64
	Text     string
	DoneChan chan int64
}

var ch chan msgChan

func InitChan() {
	ch = make(chan msgChan, 10)
}

func Send2Chan(msg msgChan) {
	ch <- msg
}

func GetFromChan() chan msgChan {
	return ch
}

// Producer
func producer(i int64) *msgChan {
	var msg msgChan
	msg.Id = i
	msg.Text = fmt.Sprintf("消息文本%v", i)
	msg.DoneChan = make(chan int64)
	Send2Chan(msg) // 发送到channel
	fmt.Println("producer: ", i)
	return &msg
}

// Consumer
func consumer() {
	for {
		select {
		case msg, ok := <-GetFromChan():
			if ok {
				fmt.Printf("consumer %v processing ..., time: %v\n",
					msg.Id, time.Now().Format("2006-01-02 15:04:05"))

				// 因为 rateLimit = 5，所以模拟每个go大概需运行多长时间
				// 也可以在外部采用time.tick的方式
				time.Sleep((1000 / rateLimit) * time.Millisecond)

				msg.DoneChan <- msg.Id // 通知生产者已经消费完了
			}
		}
	}
}

func main() {
	InitChan()

	// todo 消费者
	go consumer()

	// todo 生产者
	var wg sync.WaitGroup
	wg.Add(20) // 确保每个go都能运行完
	for i := 0; i < 20; i++ {
		msg := producer(int64(i))
		// todo 等待消费者消费完的通知（哪个先消费完就接收哪个）
		go func(m *msgChan, w *sync.WaitGroup) {
			defer w.Done()
			if v, ok := <-m.DoneChan; ok {
				fmt.Println("receive done: ", v)
				close(m.DoneChan) // 通信完后关闭channel
			}
		}(msg, &wg)
	}
	wg.Wait()
	close(ch) // 结束的时候关闭channel
	fmt.Println("后续操作...")
}
