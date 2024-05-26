package main

import (
	"testing"
)

// 测试抽奖算法
func TestDraw(t *testing.T) {
	arr := []int{20, 34, 160, 2}

	// 用 map 记录中奖id和对应次数
	winnerCount := make(map[int]int)
	// 抽奖次数
	trials := 100000

	for i := 0; i < trials; i++ {
		winner := Draw(arr)
		winnerCount[winner]++
	}

	// 遍历出每个用户的中奖次数
	for id, count := range winnerCount {
		t.Logf("User ID %d won %d times", id, count)
	}

	if winnerCount[0] == 0 || winnerCount[1] == 0 || winnerCount[2] == 0 || winnerCount[3] == 0 {
		t.Errorf("One of the users never won, which is highly unlikely")
	}
}

// 单个用户进行抽奖
func TestDrawSingleUser(t *testing.T) {
	arr := []int{100}
	winner := Draw(arr)
	if winner != 0 {
		t.Errorf("Expected winner to be 0, got %d", winner)
	}
}

// 同分数进行抽奖
func TestDrawEqualPoints(t *testing.T) {
	arr := []int{50, 50, 50, 50}

	winnerCount := make(map[int]int)
	trials := 100000

	for i := 0; i < trials; i++ {
		winner := Draw(arr)
		winnerCount[winner]++
	}

	// 遍历出每个用户的中奖次数
	for id, count := range winnerCount {
		t.Logf("User ID %d won %d times", id, count)
	}

	if len(winnerCount) != 4 {
		t.Errorf("Expected all users to win at least once")
	}
}

// 大量用户进行抽奖
func TestDrawLargeUser(t *testing.T) {
	arr := make([]int, 50000)
	for i := range arr {
		arr[i] = 1
	}

	winner := Draw(arr)
	if winner < 0 || winner >= 50000 {
		t.Errorf("Winner out of bounds: %d", winner)
	}
}
