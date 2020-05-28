package basic

import (
	"fmt"
	"testing"
)

func Test_Clist(t *testing.T) {
	c := NewClist(&Node{data: 1})
	c.InsertAfter(c.tail, 2)
	c.InsertAfter(c.tail, 3)
	c.InsertAfter(c.tail, 4)

	fmt.Println(c.head.Data())//1
	fmt.Println(c.head.next.Data())//2
	fmt.Println(c.head.next.next.Data())//3
	fmt.Println(c.head.next.next.next.Data())//4
	fmt.Println(c.head.next.next.next.next.Data())//1
}
