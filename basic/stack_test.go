package basic

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	s := NewStack(&Node{data: 1})

	s.Push(2)
	s.Push(3)
	fmt.Println(s.Top()) //3

	fmt.Println(s.PopValue()) //3

	s.Push(4)
	fmt.Println(s.PopValue()) //4
	fmt.Println(s.PopValue()) //2
	fmt.Println(s.Size())     //1
}
