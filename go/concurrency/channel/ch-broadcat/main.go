package main

import (
	"fmt"
	"sync"
)

// link: https://juejin.cn/post/6844903857395335182

type BroadcastService struct {
	// 消费者监听的广播通道
	chBroadcast chan int
	// 转发给这些 observers的通道
	chListeners []chan int
	// add observer
	chNewRequests chan (chan int)
	// remove observer
	chRemoveRequests chan (chan int)
}

// create Broadcast service
func NewBroadcastService() *BroadcastService {
	return &BroadcastService{
		chBroadcast:      make(chan int),
		chListeners:      make([]chan int, 3),
		chNewRequests:    make(chan chan int),
		chRemoveRequests: make(chan chan int),
	}
}

// create new observer and return listen channal
// 这会创建一个新消费者并返回一个监听通道
func (bs *BroadcastService) Listener() chan int {
	ch := make(chan int)
	bs.chNewRequests <- ch
	return ch
}

// remove a observer
func (bs *BroadcastService) RemoveListener(ch chan int) {
	bs.chRemoveRequests <- ch
}

func (bs *BroadcastService) addListener(ch chan int) {
	for i, v := range bs.chListeners {
		if v == nil {
			bs.chListeners[i] = ch
			return
		}
	}
	bs.chListeners = append(bs.chListeners, ch)
}

func (bs *BroadcastService) removeListener(ch chan int) {
	for i, v := range bs.chListeners {
		if v == ch {
			bs.chListeners[i] = nil
			// 要关闭 否则监听他的Goroutine会一直block
			close(ch)
			return
		}
	}
}

func (bs *BroadcastService) Run() chan int {
	go func() {
		for {
			// 处理新建消费者或移除消费者
			select {
			case newCh := <-bs.chNewRequests:
				bs.addListener(newCh)
			case removeCh := <-bs.chRemoveRequests:
				bs.removeListener(removeCh)
			case v, ok := <-bs.chBroadcast:
				// 如果广播广播channel关闭，则关闭所有的所有的observer通道
				if !ok {
					goto terminate
				}

				// 将值转发到所有的 observer(消费者) channel
				for _, dstCh := range bs.chListeners {
					if dstCh == nil {
						continue
					}
					dstCh <- v
				}

			}
		}
	terminate:
		// 关闭所有的 observer 通道
		for _, dstCh := range bs.chListeners {
			if dstCh == nil {
				continue
			}
			close(dstCh)
		}
	}()
	return bs.chBroadcast
}

func main() {
	var wg sync.WaitGroup

	bs := NewBroadcastService()
	// Broadcast publish
	chBroadcast := bs.Run()
	// Observer A
	chA := bs.Listener()
	// Observer B
	chB := bs.Listener()
	wg.Add(2)
	go func() {
		for v := range chA {
			fmt.Println("A", v)
		}
		wg.Done()
	}()
	go func() {
		for v := range chB {
			fmt.Println("B", v)
		}
		wg.Done()
	}()

	for i := 0; i < 3; i++ {
		chBroadcast <- i
	}
	close(chBroadcast)
	wg.Wait()
}
