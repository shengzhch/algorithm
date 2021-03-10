package main

import "fmt"

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？



示例 1：
输入：nums = [1,2,0] 输出：3

示例 2：
输入：nums = [3,4,-1,1] 输出：2
*/

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1})) // 2
}


func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 对数组进行归位预处理(类似于计数排序，值为7放在第七个位置 index为6)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			swap(&nums[i], &nums[nums[i]-1])
		}
	}
	// fmt.Println(nums)	// [1 -1 3 4]

	// 向后检查
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 理想正整数数组
	return n + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
