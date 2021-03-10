package main

import "fmt"

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

 

示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

*/

func main() {
	nums1 := []int{2}
	nums2 := []int{}
	merge88(nums1, 1, nums2, 0)
	fmt.Println(nums1)
}

// 两个指针指向两个数组，分别向后遍历
// 注意解决未遍历完的数组
func merge88(nums1 []int, m int, nums2 []int, n int) {
	cpNums := make([]int, 0, m+n)
	copy(cpNums, nums1)

	p1, p2 := 0, 0
	nums := nums1[:m]
	for p1 < m && p2 < n {
		if nums[p1] < nums2[p2] {
			cpNums = append(cpNums, nums[p1])
			p1++
		} else {
			cpNums = append(cpNums, nums2[p2])
			p2++
		}
	}

	if p1 < m {
		cpNums = append(cpNums, nums[p1:]...)
	}
	if p2 < n {
		cpNums = append(cpNums, nums2[p2:]...)
	}

	for i, num := range cpNums {
		nums1[i] = num // 函数内部修改 slice 只能之指定索引修改，与指定传指针原理相同
	}
}
