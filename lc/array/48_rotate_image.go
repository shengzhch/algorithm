package main

import "fmt"

/*
n阶矩阵顺时针旋转90度
*/
func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	matrix = rotate222(matrix)
	for _, nums := range matrix {
		for _, num := range nums {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}

// 正方形任意旋转可改为多重折叠
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	if n <= 1 {
		return matrix
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 左右折叠
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	return matrix
}

// 正方形任意旋转可改为多重折叠
func rotate222(matrix [][]int) [][]int {
	n := len(matrix)
	if n <= 1 {
		return matrix
	}
	// 反对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}

	//// 上下折叠
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	return matrix
}
