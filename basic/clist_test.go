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
	c := NewClist()
	c.InsertAsTail(n0)
	c.InsertAsTail(n1)
	c.InsertAsTail(n2)
	c.InsertAsTail(n3)
	c.InsertAsTail(n4)

	fmt.Println("当前")
	//0 1 2 3 4
	c.Traverse(PrintNode)
	fmt.Println(c.head.Data())                          //0
	fmt.Println(c.head.next.Data())                     //1
	fmt.Println(c.head.next.next.Data())                //2
	fmt.Println(c.head.next.next.next.Data())           //3
	fmt.Println(c.head.next.next.next.next.Data())      //4
	fmt.Println(c.head.next.next.next.next.next.Data()) //5

	fmt.Println("删除n0 - 1 2 3 4")
	c.Delete(n0)

	c.Traverse(PrintNode)

	fmt.Println("删除n4 - 1 2 3")
	c.Delete(n4)
	c.Traverse(PrintNode)

	fmt.Println("删除n2 - 1 3")
	c.Delete(n2)
	c.Traverse(PrintNode)
	fmt.Println(c.head == n1, c.tail == n3)
	c.Delete(n1)
	fmt.Println(c.head == n3, c.tail == n3)
	c.Delete(n3)

	fmt.Println(c.len, c.head, c.tail)
}
