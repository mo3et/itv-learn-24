package main

import (
	"fmt"
	"log"
	"sync"
)

func Go(x func()) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("%v\n", e)
			}
			wg.Done()
		}()
		x() // 用于传入令其 Panic 的函数，例如主动调用
	}()
	wg.Wait()
}

func main() {
	Go(tryPanic)
	Go(noPanic)
	fmt.Println("coming")
}

func tryPanic() {
	panic("干他")
}

func noPanic() {
	fmt.Println("Don't Panic")
}
