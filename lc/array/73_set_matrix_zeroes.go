package main

import "fmt"


/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
*/
func main() {
	matrix := [][]int{
		// {1, 1, 1},
		// {1, 0, 1},
		// {1, 1, 1},

		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)

	fmt.Println(matrix) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}

func setZeroes(matrix [][]int) {
	rows := make(map[int]int)
	columns := make(map[int]int)
	for i := range matrix {
		for j, num := range matrix[i] {
			if num == 0 {
				rows[i] ++
				columns[j]++
			}
		}
	}
	for r := range rows {
		for c := range matrix[0] {
			matrix[r][c] = 0
		}
	}
	for c := range columns {
		for r := range matrix {
			matrix[r][c] = 0
		}
	}
}
