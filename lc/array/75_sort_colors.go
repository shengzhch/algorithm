package main

import "fmt"

func main() {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	fmt.Println(colors)
	colors2 := []int{2, 0, 2, 1, 1, 0, 2, 0, 2, 1, 1, 0}
	sortColorsSwap(colors2)
	fmt.Println(colors2)
}

// 计数...遍历，计数不要一上来就是哈希表，把问题搞复杂了
func sortColors(nums []int) {
	zero, one, two := 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			zero++
		case 1:
			one++
		case 2:
			two++
		}
	}
	for i := 0; i < zero; i++ {
		nums[i] = 0
	}
	for i := zero; i < zero+one; i++ {
		nums[i] = 1
	}
	for i := zero + one; i < len(nums); i++ {
		nums[i] = 2
	}
}

func sortColorsSwap(nums []int) {
	n := len(nums)
	var index0 = 0
	var index2 = n - 1
	for i := 0; i < index2; {
		switch nums[i] {
		case 0:
			nums[i], nums[index0] = nums[index0], nums[i]
			index0++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[index2] = nums[index2], nums[i]
			index2--
		}
	}
}
