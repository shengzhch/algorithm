package main

import "fmt"

/*
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
请你在旋转后的数组中搜索 target，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

思路，旋转后的数据局部有序，中间节点分开的两部分肯定有一部分是有序的。



示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0 输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3 输出：-1
*/

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
}

// 类二分搜索
// 最左边数 < 中间数则左侧有序，最右边数 > 中间数则右侧有序
// 在缩小搜索区域时，一直只在确定的有序区域内查找
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case nums[mid] == target: // bingo
			return mid
		case nums[l] <= nums[mid]: // 左侧有序(从数组中间靠右的某个点旋转)
			if nums[l] <= target && target < nums[mid] { // 保证 target 一定在有序的左侧内
				r = mid - 1
			} else { //不在l～mid中，修改l的值，直接赋值l为mid+1
				l = mid + 1
			}
		case nums[mid] <= nums[r]: // 右侧有序
			if nums[mid] < target && target <= nums[r] { // 保证 target 一定在有序的右侧内
				l = mid + 1
			} else { //不在mid～r中，修改r的值，直接赋值r为mid-1
				r = mid - 1
			}
		}
	}
	return -1
}
