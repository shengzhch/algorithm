package basic

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	println("-----")
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	//3 2 1
	s.Traverse(PrintNode)
	println("-----")
	//3
	fmt.Println(s.Top())
	//3
	fmt.Println(s.PopValue()) //3
	println("-----")

	//4 2 1
	s.Push(4)
	s.Traverse(PrintNode)
	println("-----")

	//4
	fmt.Println(s.PopValue())
	//2
	fmt.Println(s.Size())
	println("-----")

	//2
	fmt.Println(s.PopValue())
	//1
	fmt.Println(s.Size())
	println("-----")
	//1
	s.Traverse(PrintNode)
	println("-----")
	//1
	fmt.Println(s.PopValue())
	//0
	fmt.Println(s.Size())
	println("-----")
	//""
	s.Traverse(PrintNode)
	println("-----")
}
