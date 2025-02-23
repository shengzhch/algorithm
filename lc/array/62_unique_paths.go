package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(7, 3)) //  3行 7列
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	var paths = make([][]int, m)

	//第一列置为1
	for r := 0; r < m; r++ {
		paths[r] = make([]int, n)
		paths[r][0] = 1
	}

	//第一行置为1
	for c := 0; c < n; c++ {
		paths[0][c] = 1
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ { // 每行从后向前走
			paths[r][c] = paths[r-1][c] + paths[r][c-1]
		}
	}

	return paths[m-1][n-1]
}
