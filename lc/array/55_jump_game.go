package main

import "fmt"

/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。
*/

func main() {
	fmt.Println(canJump2([]int{0, 3, 1, 0, 4, 3, 2, 0, 1}))

	fmt.Println(jump([]int{2, 3, 1, 1, 4, 0, 4}, true)) // true
	fmt.Println(jump([]int{2, 3, 1, 1, 4}, true))       // true
	fmt.Println(jump([]int{3, 2, 1, 0, 4}, true))       // false
	fmt.Println(jump([]int{1, 2, 3}, true))             // true
	fmt.Println(jump([]int{0, 1}, true))                // false
	fmt.Println(jump([]int{0, 2, 3}, true))             // false
	fmt.Println(jump([]int{2, 0, 0}, true))             // true
	fmt.Println(jump([]int{2, 0, 1, 0, 1}, true))       // false
}

func canJump(nums []int) bool {
	return jump(nums, true)
}

// 要想达到最后一位，数组的前 n-1 个成员必须 >= 直达数组的各成员
// 如 [2 3 1 1] >=  [4 3 2 1]
//    [2, 3, 1] >= [3 2 1]
//    [2, 3] >= [2 1]
//    [2] >= [1]

// 效率比较低的递归
func jump(nums []int, ok bool) bool {
	n := len(nums)
	if !ok {
		return false
	}

	// 处理遇到 0 的极端情况
	for i := n - 1 - 1; i >= 0; i-- {
		if nums[i] == 0 {
			direct := 1
			for j := i - 1; j >= 0; j-- {
				// fmt.Println(nums[j], direct)
				//只要能迈过0就能到达队尾，后面都是正数，一定可达
				if nums[j] > direct {
					return jump(nums[:i], true)
				}
				direct++
			}
			return false
		}
	}
	return true
}

func canJump2(nums []int) bool {
	n := len(nums)

	for i := n - 1 - 1; i >= 0; i-- {
		if nums[i] == 0 {
			direct := 1
			for j := i - 1; j >= 0; j-- {
				if nums[j] > direct {
					fmt.Println("nums[:j] ", nums[:j])
					return canJump2(nums[:i])
				}
				direct++
			}
			return false
		}
	}
	return true
}
