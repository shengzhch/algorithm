package main

import "fmt"

/*

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1：

输入：nums = [3,2,2,3], val = 3 输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，
而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

*/
func main() {
	fmt.Println(removeElement1([]int{3, 2, 2, 4, 1, 3}, 3))
}

//
// 遍历数组一如既往地联想到双指针法
//
func removeElement1(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast] // 快指针在前，慢指针在后，遇到不等的元素就放到慢指针的位置
			slow++
		}
	}
	fmt.Println(nums[:slow])
	return slow
}
