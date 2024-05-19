package main

import (
	"fmt"
	"sync"
)

var (
	numChan, strChan = make(chan struct{}), make(chan struct{})
	waitg            sync.WaitGroup
)

// print number
func printNum() {
	defer waitg.Done()
	for i := 1; i < 27; i += 2 {
		<-numChan                  // 阻塞直到字幕被打印后，numChan 写入
		fmt.Printf("%v%v", i, i+1) // 打印数字
		strChan <- struct{}{}      // strChan写入，打印字母的协程的strChan取出，才会打印字母
	}
	<-numChan // 读取最后一个从Goroutine2传入的，防止死锁
}

// print char
func printStr() {
	defer waitg.Done()
	for i := 65; i < 91; i += 2 {
		<-strChan
		fmt.Printf("%v%v", string(byte(i)), string(byte(i+1)))
		numChan <- struct{}{}
	}
}

func main() {
	waitg.Add(2)
	go printNum()
	go printStr()
	numChan <- struct{}{}
	waitg.Wait()
}
