package main

import "fmt"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1] 输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*/

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
}

func trap(height []int) int {
	area := 0
	l, r := 0, len(height)-1
	lTop, rTop := 0, 0

	//最小桶理论
	for l <= r {
		lTop = max(height[l], lTop)
		rTop = max(height[r], rTop)
		if lTop < rTop { //右边高，限制在左边
			area += lTop - height[l] // 右侧更高些，往右侧走。一边向右移，一边计算高度差来累计面积（最高点处高度差为 0）
			l++
		} else { //左边高，限制在右边
			area += rTop - height[r] // 向左侧移，用目前右侧最高高度来计算高度差，累计面积
			r--
		}
	}
	return area
}
