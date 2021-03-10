package main

import "fmt"

/*

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4] 输出：6 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。


示例 2：
输入：nums = [1] 输出：1
*/

func main() {
	fmt.Println(maxSubArray([]int{-3, 4, -1, 3, -5}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}

//
// 跨步遍历的感觉
//
func maxSubArray(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	sum, maxSum := 0, 0
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] > 0:
			sum += nums[i]
		case nums[i] < 0:
			if sum+nums[i] > 0 {
				sum += nums[i]
			} else {
				sum = 0
			}
		}

		if maxSum < sum {
			maxSum = sum
		}
	}

	//maxSum为0，从未修改过
	if max < 0 && maxSum > max {
		return max
	}
	return maxSum
}
