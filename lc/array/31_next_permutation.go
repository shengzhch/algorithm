// 参考 https://goleetcode.io/2018/11/20/array/31-next-permutation
package main

import "fmt"

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须 原地 修改，只允许使用额外常数空间。



示例 1：
输入：nums = [1,2,3] 输出：[1,3,2]

示例 2：
输入：nums = [3,2,1] 输出：[1,2,3]

示例 3：
输入：nums = [1,1,5] 输出：[1,5,1]

示例 4：
输入：nums = [1] 输出：[1]



提示：
1 <= nums.length <= 100
0 <= nums[i] <= 100
*/

func main() {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums) // [1 3 1 2 4 7] // bingo
}

// 数组规律题
//字典顺序要求数组改动尽可能的靠后
// 从后往前找第一个下降点 i，再从后往前找它的 ceil 值，交换
// 再将 [i+1:] 之后的数据从降序反转为升序，为最小序列
func nextPermutation(nums []int) {

	// 处理降序的 case
	desc := true
	n := len(nums)
	for i := range nums[:n-1] {
		if nums[i] < nums[i+1] {
			desc = false
		}
	}
	if desc {
		reverse(nums)
		return
	}

	// 从后向前找第一个下降的点，（找到后替换该点，用之后的第一个大于该点的值替换）
	var i int
	for i = n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			i-- // 找到 2
			break
		}
	}

	// 从后向前，找向上最接近的值
	for j := n - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j] // 交换 2 和 3 	// [1 3 7 4 2 1]
			break
		}
	}

	//i+1之后的数组之前是降序，然后改为升序
	reverse(nums[i+1:]) // 反转 4 2 1	// [1 3 1 2 4 7]
}
