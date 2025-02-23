package main

import (
	"algorithm/basic"
	"fmt"
)

//部分反转

func Reversepart(l *basic.List, from, to int) {
	l.ReversePart(from, to)
}

func main() {
	l1 := basic.NewListWithData(1, 2, 3, 4, 5, 6)
	Reversepart(l1, 1, 2)
	l1.Traverse(basic.PrintNode)

	fmt.Println("-------")

	l1.Reverse()
	l1.Traverse(basic.PrintNode)

	fmt.Println("-------")
	l1.ReverseAsSingle()
	l1.Traverse(basic.PrintNode)

	fmt.Println("-------")
	Reversepart(l1, 1, 2)
	l1.Traverse(basic.PrintNode)
}
