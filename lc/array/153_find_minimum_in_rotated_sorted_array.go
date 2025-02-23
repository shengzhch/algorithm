package main

import "fmt"

//寻找旋转排序数组中的最小值,旋转的意思是某个位置之后的元素全部放到数组头部

//思路 目标值的位置一定在无序的部分



func main() {
	// 在数组中查找原本 3 该在的位置，即是发生旋转的值（最小值）
	fmt.Println(findMin([]int{4, 5, 1, 2, 3}))
}

func findMin(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}

	// 若在最后一个元素上旋转，数组依旧有序
	if nums[n-1] > nums[0] {
		return nums[0]
	}


	target := nums[0]
	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		// 不断在无序区找最小值
		switch {
		case nums[mid] >= target: // 中间值比目标大，前半部分有序，继续在后部分无序区找最小值
			l = mid + 1
		case nums[mid] < target: // 中间值比目标小，前半部分无序
			r = mid
		}
	}
	return nums[r]
}
