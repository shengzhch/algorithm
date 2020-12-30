package basic

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
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

	//0 1 2 3 4 5
	fmt.Println("当前")
	l.Traverse(PrintNode)

	//5 4 3 2 1 0
	fmt.Println("反转后")
	l.Reverse()
	l.Traverse(PrintNode)

	//0 1 2 3 4 5
	l.Reverse()

	//1 2 3 4 5
	fmt.Println("删除头节点")
	l.DelHead()
	l.Traverse(PrintNode)

	//1 2 3 4
	fmt.Println("删除尾节点")
	l.DelTail()
	l.Traverse(PrintNode)

	//1 2 2.5 3 4
	fmt.Println("插入2.5")
	n25 := &Node{data: 2.5}
	l.InsertAfter(n2, n25)
	l.Traverse(PrintNode)

	//1 1.5 2 2.5 3 4
	fmt.Println("插入1.5")
	n15 := &Node{data: 1.5}
	l.InsertBefore(n2, n15)
	l.Traverse(PrintNode)

	//1 1.5 2.5 3 4
	fmt.Println("删除2")
	l.DelNode(n2)
	l.Traverse(PrintNode)

	//4 3 2.5 1.5 1
	fmt.Println("反转后")
	l.Reverse()
	l.Traverse(PrintNode)

	//1 1.5 2.5 3 4
	fmt.Println("单链表反转后")
	l.ReverseAsSingle()
	l.Traverse(PrintNode)

	//true true
	fmt.Println(l.head == n1)
	fmt.Println(l.tail == n4)
	fmt.Println(l.len == 5)
}
