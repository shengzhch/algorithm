package main

import (
	"fmt"
)

/*
给你一个整数数组 nums ，数组中的元素互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 6}))
	fmt.Println("-----------")
	fmt.Println(subsets2([]int{9, 0, 3, 5, 6}))
	//fmt.Println(subsets2([]int{1, 2, 3, 4}))
}

// 自创的 4ms 解法阿西吧
// 理解子集的本质，N 个数有 1~N 种组合方式
// 可以直接从后向前遍历，每遇到一个数就把它与所有的现有子集重新 append 成新子集
// 遍历完毕就是所有子集
// 我觉得 ok
func subsets(nums []int) [][]int {
	n := len(nums)
	seqs := make([][]int, 1)
	seqs = append(seqs, []int{nums[n-1]})

	for i := n - 2; i >= 0; i-- {
		for _, seq := range seqs {
			nextSeq := append([]int{nums[i]}, seq...)
			seqs = append(seqs, nextSeq)
		}
	}

	return seqs
}

func subsets2(nums []int) [][]int {
	n := len(nums)
	seqs := make([][]int, 1)

	for i := 0; i < n; i++ {
		for _, seq := range seqs {
			nextSeq := append(clone(seq), nums[i])
			seqs = append(seqs, nextSeq)
		}
	}

	return seqs
}

//注意数组append的坑
func clone(src []int) []int {
	var dst = make([]int, len(src), len(src))
	copy(dst, src)
	return dst
}
