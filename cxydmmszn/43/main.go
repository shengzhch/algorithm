package main

import (
	"algorithm/basic"
	"fmt"
)

func josephusWithCList(l *basic.List, m int) []interface{} {

	rel := make([]interface{}, 0, l.Size()-1)
	if l.Size() <= 1 {
		return []interface{}{l.Head().GetValue()}
	}

	cur := l.Head()

	var pre *basic.ListElm

	count := 0

	for l.Size() != 1 {
		count ++
		if count == m {
			rel = append(rel, cur.GetValue())
			l.Rem_next(pre)
			count = 0
		}
		pre = cur
		cur = cur.Next()
	}
	return rel
}

func josephusNodes(l *basic.ListElm, m int) []interface{} {
	//todo
	return nil
}

func main() {

	l1 := &basic.List{}
	(*basic.CList)(l1).Ins_next((*basic.List)(l1).Tail(), 1)
	(*basic.CList)(l1).Ins_next((*basic.List)(l1).Tail(), 2)
	(*basic.CList)(l1).Ins_next((*basic.List)(l1).Tail(), 3)
	(*basic.CList)(l1).Ins_next((*basic.List)(l1).Tail(), 4)
	(*basic.CList)(l1).Ins_next((*basic.List)(l1).Tail(), 5)

	//l1.Traverse(basic.Tf)

	rel := josephusWithCList(l1, 2)

	fmt.Println("rel ", rel) //2415

}
