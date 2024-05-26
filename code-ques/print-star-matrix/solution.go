package main

import "fmt"

// ● 给定一个字符的矩形矩阵，为其添加带星号的边框。
// "abc"
// *****
// *abc*
// *****

func solution(matrix []string) []string {
	if len(matrix) == 0 {
		return []string{}
	}
	width := len(matrix[0])

	var borderedMatrix []string

	border := ""
	for i := 0; i < width+2; i++ {
		border += "*"
	}

	// 添加顶部
	borderedMatrix = append(borderedMatrix, border)

	for _, row := range matrix {
		borderedRow := "*" + row + "*"
		borderedMatrix = append(borderedMatrix, borderedRow)
	}

	// 添加底部
	borderedMatrix = append(borderedMatrix, border)
	return borderedMatrix
}

func printMartix(matrix []string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
