package main

import "fmt"

func main() {
	fmt.Println(maxProfit1([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit1(prices []int) int {
	i, max := 0, 0
	n := len(prices)
	for i < n-1 {
		buy := i
		for i < n-1 && prices[i] < prices[i+1] {
			i++
		}
		max += prices[i] - prices[buy] // 在最高点抛售
		i++
	}
	return max
}

func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
