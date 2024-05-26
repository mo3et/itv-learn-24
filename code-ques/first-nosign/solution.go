package main

import "sort"

// 未签到的第一人
// 根据已签到的n个编号，计算未签到的人的最小编号是多少

// 输入 sign={4,2,1,6,5,7} n=6
// 输出 3

func solution(sign []int) int {
	sort.Ints(sign)

	missing := 1
	for _, num := range sign {
		if num == missing {
			missing++
		} else if num > missing {
			break
		}
	}
	return missing
}

func main() {
}
