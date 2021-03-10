package main

import (
	"algorithm/basic"
	"fmt"
)

func revarseList(l *basic.List) {
	l.ReverseAsSingle()

}

func revarseDList(l *basic.List) {
	l.Reverse()
}

func main() {

	l1 := basic.NewListWithData(1, 2, 3, 4, 5, 6)
	revarseDList(l1)
	l1.Traverse(basic.PrintNode)

	fmt.Println("-----------")

	revarseList(l1)
	l1.Traverse(basic.PrintNode)
}
