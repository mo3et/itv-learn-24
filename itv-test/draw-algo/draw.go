package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Draw algo
// 只用两个独立的 range 进行计算和查找，时间复杂度为O(n)
// 范围 1 <=len(arr)<=50000 和 1<=arr[i]<=50000
func Draw(arr []int) int {
	// 超过范围
	if len(arr) < 1 || len(arr) > 50000 {
		log.Fatal("人数不符合规范")
	}
	// 计算积分总和
	sum := 0
	for _, score := range arr {
		if score < 1 || score > 50000 {
			log.Fatal("分数不符合规范")
		}
		sum += score
	}
	// 生成随机数
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	randNum := r.Intn(sum)

	// 查找中奖用户
	for i, score := range arr {
		randNum -= score
		if randNum < 0 {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{20, 24, 160, 2}
	winner := Draw(arr)
	fmt.Printf("中奖者ID是:%d\n", winner)
}
