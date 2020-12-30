package basic

import (
	"fmt"
	"testing"
)

func Test_Clist(t *testing.T) {

	n0 := &Node{data: 0}
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}
	n4 := &Node{data: 4}
	c := NewClist(n0)
	c.InsertAfter(c.tail, n1)
	c.InsertAfter(c.tail, n2)
	c.InsertAfter(c.tail, n3)
	c.InsertAfter(c.tail, n4)

	fmt.Println("当前")
	//0 1 2 3 4
	c.Traverse(PrintNode)
	fmt.Println(c.head.Data())                          //0
	fmt.Println(c.head.next.Data())                     //1
	fmt.Println(c.head.next.next.Data())                //2
	fmt.Println(c.head.next.next.next.Data())           //3
	fmt.Println(c.head.next.next.next.next.Data())      //4
	fmt.Println(c.head.next.next.next.next.next.Data()) //5

	fmt.Println("删除n0的下一个节点 0 2 3 4")
	c.DeleteAfter(n0)

	c.Traverse(PrintNode)

	fmt.Println("删除n4的下一个节点 2 3 4")
	c.DeleteAfter(n4)
	c.Traverse(PrintNode)

	fmt.Println("删除n3的下一个节点 2 3")
	c.DeleteAfter(n3)
	c.Traverse(PrintNode)
}
