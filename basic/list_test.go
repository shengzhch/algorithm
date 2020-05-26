package basic

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	l := new(List)
	n0 := &Node{data: 0}
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}
	n4 := &Node{data: 4}
	n5 := &Node{data: 5}

	l.Init(n0)
	l.InsertAsTail(n1)
	l.InsertAsTail(n2)
	l.InsertAsTail(n3)
	l.InsertAsTail(n4)
	l.InsertAsTail(n5)

	fmt.Println("当前")
	l.Traverse(PrintNode)

	fmt.Println("反转后")
	l.Reverse()
	l.Traverse(PrintNode)

	l.Reverse()

	fmt.Println("删除头节点")
	l.DelHead()
	l.Traverse(PrintNode)

	fmt.Println("删除尾节点")
	l.DelTail()
	l.Traverse(PrintNode)

	fmt.Println("插入2.5")
	n25 := &Node{data: 2.5}
	l.InsertAfter(n2, n25)
	l.Traverse(PrintNode)

	fmt.Println("插入1.5")
	n15 := &Node{data: 1.5}
	l.InsertBefore(n2, n15)
	l.Traverse(PrintNode)

	fmt.Println("删除2")
	l.DelNode(n2)
	l.Traverse(PrintNode)

	fmt.Println("反转后")
	l.Reverse()
	l.Traverse(PrintNode)

	fmt.Println("单链表反转后")
	l.ReverseAsSingle()
	l.Traverse(PrintNode)
}
