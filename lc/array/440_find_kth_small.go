package main

import "fmt"

//给定整数n和k，找到1-n中字典树中第k小的数，其中1<=k<=n<10^9

func main() {
	fmt.Println(finckthNum(13, 3))
}

//计算某个前缀下各有多少个子节点(思路：下一个前缀的起点减去当前前缀的起点)
func getCount(prefix int, n int) int {
	count := 0
	cur := prefix
	next := prefix + 1

	for cur <= n {
		count = count + int(min(next, n+1)) - cur
		cur = cur * 10
		next = next * 10
	}
	return count
}

func finckthNum(n int, k int) int {
	curPre := 1   //当前指向的数字
	curIndex := 1 //当前数字的位置

	for curIndex < k {
		count := getCount(curPre, n)
		//在该前缀下,第k个数在该前缀下,必须严格大于，记录count的是包含了pre本身
		if curIndex+count > k {
			curPre = curPre * 10
			//count = count + next - cur		curIndex++
		} else {
			curPre++
			curIndex = curIndex + count
		}
	}

	return curPre
}
