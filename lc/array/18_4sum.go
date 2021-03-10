package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
	fmt.Println(fourSum([]int{-3, 0, 0, 1, 2,}, 0))    // [[-3 0 1 2]]
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))         // [[0 0 0 0]]
}

/*

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。 满足要求的四元组集合为： [ [-1, 0, 0, 1], [-2, -1, 1, 2], [-2, 0, 0, 2] ]

*/

//
// N 数之和的本质是，在有序数组内，寻找 N 个数的和恰好是 S
// 解决办法还是 3sum 3sum_closest 的双指针法，不过需要外部 N-2 层循环，内部双指针循环即可
// 注意双指针在遍历时外部所有循环要去重，指针移动时也要去重
//
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-1; i++ {
		//前面的接口肯定包含后面的结果
		if i > 0 && nums[i] == nums[i-1] { // 去重1
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] { // 去重2	// 注意条件：j>i+1 与 i>0 相同都是为了排除第一个相同数
				continue
			}
			head, tail := j+1, n-1
			for head < tail {
				sum := nums[i] + nums[j] + nums[head] + nums[tail]
				switch {
				case sum < target: // 向后走
					head++
				case sum > target: // 向前走
					tail--
				case sum == target: // 向前向后走
					res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
					// 去重3：注意 for 循环条件的判断，避开死循环
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					head++
					tail--
				}
			}
		}
	}
	return res
}
