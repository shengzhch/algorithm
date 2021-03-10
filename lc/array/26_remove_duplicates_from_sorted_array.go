package main

import "fmt"

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



示例 1:

给定数组 nums = [1,1,2], 函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。 你不需要考虑数组中超出新长度后面的元素。

示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4], 函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。 你不需要考虑数组中超出新长度后面的元素。
*/
func main() {
	fmt.Println(removeDuplicates1([]int{1, 1, 2}))                      // 2
	fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) // 5
}

// 针对有序数组，双指针法是十分常见且有用的
func removeDuplicates1(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != nums[fast+1] { // 相邻的数不相等
			slow++
			fast++
			nums[slow] = nums[fast] // 将最新的新数存储到慢指针的位置
			continue
		}
		fast++
	}
	fmt.Println(nums[:slow+1])
	return slow + 1
}

// 充分利用数组有序的已知条件,思想还是双指针法
func removeDuplicates2(nums []int) int {
	n := len(nums)
	l, r := 0, 1
	for r < n {
		if nums[l] < nums[r] { // 比我大就放到我的下一个，并用于下一次与r的比较
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
		r++
	}
	fmt.Println(nums[:l+1])
	return l + 1
}
