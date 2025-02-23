package main

import (
	"algorithm/basic"
	"fmt"
)

//在单链表和双链表删除倒数第K个节点

func deleteListKNode(l *basic.List, k int) {
	head := l.Head()
	if head == nil || k < 1 {
		return
	}
	cur := head
	for cur != nil {
		k--
		cur = cur.Next()
	}
	if k > 0 {
		//不存在
		return
	}
	if k == 0 {
		l.DelHead()
		return
	}

	if k < 0 {
		cur = head
		for k != 0 {
			cur = cur.Next()
			k++

		}
		l.DelNode(cur)
		return
	}
}

func main() {
	l1 := basic.NewListWithData(1, 2, 3, 4, 5)

	deleteListKNode(l1, 5)

	fmt.Println("删除1 结果期望 2 3 4 5")

	l1.Traverse(basic.PrintNode)

	deleteListKNode(l1, 1)

	fmt.Println("删除5 结果期望 2 3 4")

	l1.Traverse(basic.PrintNode)
}
