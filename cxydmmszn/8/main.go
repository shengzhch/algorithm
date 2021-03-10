package main

import (
	"algorithm/basic"
	"fmt"
)

// 递归调用取出插入,最终取出栈底元素
func getAndRemoveLastEle(s *basic.Stack) interface{} {
	rel := s.PopValue()
	if s.Size() == 0 {
		return rel
	} else {
		last := getAndRemoveLastEle(s)
		s.Push(rel)
		return last
	}
}

//递归调用反转插入,最终实现stack的逆序
func ReverseStack(s *basic.Stack) {
	if s.Size() == 0 {
		return
	} else {
		data := getAndRemoveLastEle(s)
		ReverseStack(s)
		s.Push(data)
	}
}

func main() {

	ms := &basic.Stack{}
	ms.Init()

	ms.Push(1)
	ms.Push(2)
	ms.Push(3)
	ms.Push(4)
	ms.Push(5)

	fmt.Println("before ------- ")
	(*basic.List)(ms).Traverse(basic.PrintNode)

	ReverseStack(ms)
	fmt.Println("after ------- ")

	(*basic.List)(ms).Traverse(basic.PrintNode)
}
