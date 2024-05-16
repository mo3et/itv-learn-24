package main

import (
	"fmt"
	"sync"
)

var (
	ch1, ch2 = make(chan struct{}, 1), make(chan struct{})
	wg       sync.WaitGroup
)

// 打印奇数
func odd() {
	defer wg.Done()
	for i := 1; i <= 100; i += 2 {
		<-ch1
		fmt.Println("go1:", i)
		ch2 <- struct{}{}
	}
}

// 打印偶数
func even() {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		<-ch2
		fmt.Println("go2:", i)
		ch1 <- struct{}{}
	}
}

func print_even_odd() {
	wg.Add(2)
	go odd()
	go even()
	ch1 <- struct{}{}
	wg.Wait()
}
