package main

import (
	"fmt"
	"math"
)

//你有三种硬币，2元、5元、7元，每种硬币足够多，买一本书需要27元，用最少的硬币组合

func coinChange(a []int, m int) int {
	var f = make([]int, m+1, m+1)
	f[0] = 0
	n := len(a)
	for i := 1; i <= m; i++ {
		f[i] = math.MaxInt16
		for j := 0; j < n; j++ {
			if i >= a[j] && f[i-a[j]] != math.MaxInt16 {
				f[i] = min(f[i-a[j]]+1, f[i])
			}
		}
	}
	return f[m]
}

func main() {
	fmt.Println(coinChange([]int{2, 5, 7}, 27))
}
