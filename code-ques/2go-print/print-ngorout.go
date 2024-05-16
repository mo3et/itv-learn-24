package main

import "fmt"

func main() {
	print_ngo()
}

func print_ngo() {
	goroutineNum := 5
	var chanSlice []chan int
	exitChan := make(chan int)

	for i := 0; i < goroutineNum; i++ { // 创建N个Goroutine
		chanSlice = append(chanSlice, make(chan int, 1))
	}

	num := 1
	j := 0
	for i := 0; i < goroutineNum; i++ {
		go func(i int) {
			for {
				<-chanSlice[i] // 循环阻塞等待
				if num > 100 {
					exitChan <- 1
					break
				}

				fmt.Println(fmt.Sprintf("goroutine%v:%v", i, num))
				num++

				if j == goroutineNum-1 { // j来控制启动那个Goroutine打印
					j = 0
				} else {
					j++
				}
				chanSlice[j] <- 1
			}
		}(i)
	}
	chanSlice[0] <- 1

	select {
	case <-exitChan:
		fmt.Println("end")
	}
}
