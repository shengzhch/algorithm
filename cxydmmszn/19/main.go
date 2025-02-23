package main

import (
	"algorithm/basic"
	"fmt"
)

func getMaxArrfromArr(arr []interface{}, w int, cf basic.CF) []interface{} {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}

	//双向链表（这里使用的双向链表，其实单链表（包含头结点和尾节点）即可）
	//存储 数组的下标
	qmax := new(basic.List)
	qmax.Init(nil)

	res := make([]interface{}, len(arr)-w+1)

	index := 0

	for i := 0; i < len(arr); i++ {

		//把小于等于当前值的元素全部移出链表
		for qmax.Len() > 0 && cf(arr[qmax.Tail().Data().(int)], arr[i]) <= 0 {
			qmax.DelNode(qmax.Tail())
		}

		//插入当前值的下标
		qmax.InsertAsTail(basic.NewNode(i))

		//第i个元素位于框中，框中最左端元素的下线i-w+1,之前的元素需要移除
		if cf(qmax.Head().Data(), i-w) == 0 {
			qmax.DelNode(qmax.Head())
		}

		//当框中元素完整是，取出此时的最大值，放入结果集中。整个赋值区间 w-1 ~ len(arr)，长度与res长度相等
		if i >= w-1 {
			res[index] = arr[qmax.Head().Data().(int)]
			index++
		}
	}

	return res
}

func getMaxArray(arr []int, w int) []int {
	rel := make([]int, 0, w)
	res := make([]int, len(arr)-w+1, len(arr)-w+1)
	var index int
	for i := 0; i < len(arr); i++ {
		for len(rel) > 0 && arr[len(rel)-1] <= arr[i] {
			rel = rel[:len(rel)-1]
		}
		rel = append(rel, i)
		if rel[0] == i-w {
			rel = rel[1:]
		}

		//满足窗口条件了，开始往res中添加元素
		if i >= w-1 {
			res[index] = arr[rel[0]]
			index++
		}
	}
	return res

}

func main() {
	arr := []interface{}{4, 3, 5, 4, 3, 3, 6, 7}
	rel := getMaxArrfromArr(arr, 3, basic.CfInt)
	fmt.Println("rel ", rel)

	arr2 := []int{4, 3, 5, 4, 3, 3, 6, 7}
	rel2 := getMaxArray(arr2, 3)
	fmt.Println("rel2 ", rel2)
}
