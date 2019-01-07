package base

import (
	"testing"
	"fmt"
)

func TestLinkList(t *testing.T) {
	l := NewLinkList()
	for i := 0; i < 10; i++ {
		l.HeadInsert((IntValue)(i))
	}

	fmt.Println(l)
	l.Clear()

	for i := 0; i < 10; i++ {
		l.RearInsert((IntValue)(i))
	}
	fmt.Println(l)
	l.Reverse()
	fmt.Println(l)


	//l.Remove((IntValue)(5))
	//fmt.Println(l)
	//l.RemoveByIndex(1)
	//fmt.Println(l)
	//l.RemoveByIndex(5)
	//fmt.Println(l)
	//l.RemoveByIndex(7)
	//fmt.Println(l)


}
