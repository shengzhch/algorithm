package main

import (
	"algorithm/basic"
	"fmt"
)

type Mystack struct {
	data *basic.Stack
	min  *basic.Stack
}

func (s *Mystack) Init() {
	s.data.Init()
	s.min.Init()
}

func (s *Mystack) Push(data interface{}, cf basic.CF) {
	if s.min.Size() == 0 {
		s.min.Push(data)
	} else if cf(data, s.Getmin()) <= 0 {
		s.min.Push(data)
	}

	s.data.Push(data)
}

func (s *Mystack) Pop(cf basic.CF) interface{} {
	pv := s.data.PopValue()
	if cf(pv, s.Getmin()) == 0 {
		s.min.PopValue()
	}
	return pv
}

func (s *Mystack) Getmin() interface{} {
	return s.min.Top()
}

func (s *Mystack) Show() {
	fmt.Println("data --- ")

	(*basic.List)(s.data).Traverse(basic.PrintNode)

	fmt.Println("min --- ")
	(*basic.List)(s.min).Traverse(basic.PrintNode)
}

func main() {

	cf := basic.CfInt
	ms := &Mystack{
		&basic.Stack{}, &basic.Stack{},
	}
	ms.Init()

	ms.Push(1, cf)
	ms.Push(2, cf)
	ms.Push(1, cf)
	ms.Push(5, cf)
	ms.Push(4, cf)
	ms.Push(3, cf)

	ms.Show()
	ms.Pop(cf)
	ms.Pop(cf)
	ms.Show()
	ms.Pop(cf)
	ms.Pop(cf)
	ms.Show()
}
