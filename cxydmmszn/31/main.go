package main

import (
	"algorithm/basic"
	"fmt"
)

func getMaxMinOfNumFromArray(arr []int, num int) int {
	if len(arr) == 0 {
		return 0
	}

	var qmin = new(basic.List)
	var qmax = new(basic.List)

	var rel int
	var i, j int

	for i < len(arr) {
		for j < len(arr) {
			for !qmax.IsEmpty() && arr[qmax.Tail().Data().(int)] <= arr[j] {
				qmax.DelTail()
			}
			qmax.InsertAsTail(basic.NewNode(j))

			for !qmin.IsEmpty() && arr[qmin.Tail().Data().(int)] >= arr[j] {
				qmin.DelTail()
			}
			qmin.InsertAsTail(basic.NewNode(j))

			if arr[qmax.Head().Data().(int)]-arr[qmin.Head().Data().(int)] > num {
				break
			}
			j++
		}

		if qmin.Head().Data().(int) == i {
			qmin.DelHead()
		}
		if qmax.Head().Data().(int) == i {
			qmax.DelHead()
		}
		fmt.Println(" min---- ")
		qmin.Traverse(basic.PrintNode) //3
		fmt.Println(" mac---- ")
		qmax.Traverse(basic.PrintNode) // 1 2 3
		fmt.Println("i,j", i, j)
		//i ... j-1; 包括arr[i,i]
		rel = rel + (j - 1 - i + 1)
		i++
	}
	return rel
}

func main() {
	rel := getMaxMinOfNumFromArray([]int{2, 4, 3, 1, 5}, 2) //9

	fmt.Println(rel)
}
