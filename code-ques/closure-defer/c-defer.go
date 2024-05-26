package main

import "fmt"

// https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html?from=groupmessage&isappinstalled=0

func main() {
	fmt.Println(DScore(0))
	fmt.Println(DScore(10.0))
	fmt.Println(DScore(20.0))
	fmt.Println(DScore(50.0))
}

func DScore(source float32) (score float32) {
	score = 2
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	score = source * 2
	return
}
