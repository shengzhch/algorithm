package main

import (
	"algorithm/basic"
	"fmt"
)

//打印两个有序链表的公共部分

func getCommonListData(l1, l2 *basic.List, cf basic.CF) []interface{} {
	s1 := l1.Head()
	s2 := l2.Head()
	rel := make([]interface{}, 0)
	for s1 != nil && s2 != nil {
		cm := cf(s1.Data(), s2.Data())
		if cm < 0 {
			s1 = s1.Next()
		} else if cm > 0 {
			s2 = s2.Next()
		} else {
			rel = append(rel, s1.Data())
			s1 = s1.Next()
			s2 = s2.Next()

		}
	}

	return rel
}

func main() {
	l1 := new(basic.List)
	l1.InsertAsTail(basic.NewNode(1))
	l1.InsertAsTail(basic.NewNode(3))
	l1.InsertAsTail(basic.NewNode(4))
	l1.InsertAsTail(basic.NewNode(6))
	l1.InsertAsTail(basic.NewNode(7))

	l2 := new(basic.List)
	l2.InsertAsTail(basic.NewNode(1))
	l2.InsertAsTail(basic.NewNode(2))
	l2.InsertAsTail(basic.NewNode(3))
	l2.InsertAsTail(basic.NewNode(5))
	l2.InsertAsTail(basic.NewNode(7))

	rel := getCommonListData(l1, l2, basic.CfInt)

	fmt.Println("rel ", rel)
}
