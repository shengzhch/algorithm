package main

import (
	"fmt"
)

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/
func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

// 考虑使用递归实现，到达 grid[i][j] 的最小路径为到达 grid[i-1][j] 和 grid[i][j-1] 的最小值
// 但递归不保存任何已计算出的结果，效率不高，类比斐波那契数列的递归实现
// 使用简单的动态规划，即保存中间计算结果
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var paths = grid

	//第一列置为1
	for r := 1; r < m; r++ {
		paths[r][0] = paths[r-1][0]+paths[r][0]
	}

	//第一行置为1
	for c := 1; c < n; c++ {
		paths[0][c] = paths[0][c-1]+paths[0][c]
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ { // 每行从后向前走
			if paths[r-1][c] <paths[r][c-1]{
				paths[r][c] = paths[r-1][c] + paths[r][c]
			}else{
				paths[r][c] = paths[r][c-1] + paths[r][c]
			}
		}
	}

	return paths[m-1][n-1]
}
