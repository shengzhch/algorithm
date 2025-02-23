package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2}))
}

// 简单的规律
func maxProfit(prices []int) int {
	n := len(prices)
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

// 记录当前股票的最低点
func maxProfit12(prices []int) int {
	var minprice = prices[0]
	var max int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > max {
			max = prices[i] - minprice
		}
	}
	return max
}

//动态规划
//d[i][0]表示这一天不持有股票的收益
//d[i][1]表示这一天持有股票的最大收益

func maxProfit13(prices []int) int {
	var rel = make([][2]int, len(prices))
	rel[0][0] = 0
	rel[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		rel[i][0] = max(rel[i-1][0], rel[i-1][1]+prices[i])
		rel[i][1] = max(rel[i-1][1], -prices[i])
	}
	return rel[len(prices)-1][0]
}

func maxProfit14(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}
