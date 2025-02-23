package main

import (
	"fmt"
	"sort"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]

*/

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combine(candidates, target)
}


//两层递归
func combine(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) <= 0 { // 递归到最后一个元素，递归结束
		return res
	}

	subSum := target - nums[0]

	//以第一个元素开头的解
	switch {
	case subSum < 0: // 向后不存在更大的值，递归结束
	    //不存在
		return res
	case subSum == 0: // 第一个数就是 target
	    //只有第一个是
		res = append(res, []int{nums[0]})
	case subSum > 0:
		//
		remains := combine(nums, subSum) // 寻找所有的子和组合
		for _, v := range remains {
			way := append([]int{nums[0]}, v...)
			res = append(res, way)
		}
	}
	//递归求其它元素开头的解
	res = append(res, combine(nums[1:], target)...) // 向后查找其他组合，避免重复
	return res
}
