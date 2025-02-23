package main

import (
	"algorithm/basic"
	"fmt"
)

//删除链表中间节点和a/b处的节点

func findk(l, a, b int) int {
	r := float64(a) / float64(b)
	if r == 0 {
		return 0
	}

	for i := 1; i <= l; i++ {
		if float64(i)/float64(l) >= r {
			fmt.Println("k ", i)
			return i
		}
	}

	return 0
}

//删除第K个节点
func deleteListKNode(l *basic.List, k int) {
	if k < 1 || k > l.Len() {
		return
	}

	var cur = l.Head()
	for i := 1; i < k; i++ {
		cur = cur.Next()
	}

	l.DelNode(cur)
}

func main() {
	l1 := basic.NewListWithData(1, 2, 3, 4, 5, 6)

	deleteListKNode(l1, findk(l1.Len(), 4, 6))

	l1.Traverse(basic.PrintNode)

}
