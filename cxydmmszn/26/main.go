package main

import (
	"algorithm/basic"
	"fmt"
)

/*
1011
1110
1111
*/

func getMaxReclFromArray(arr [][]int) int {
	if len(arr) == 0 {
		return 0
	}
	var h = make([]int, len(arr[0]))
	var maxArea int

	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[0]); col++ {
			if arr[row][col] == 0 {
				h[col] = 0
			} else {
				fmt.Println("row ,col", row, col)
				h[col] = h[col] + 1
			}
		}

		newMaxArea := getMaxArea(h)
		if newMaxArea > maxArea {
			maxArea = newMaxArea
		}
	}
	return maxArea
}

//在栈中找到第一个小于h[i]的位置，然后插入
//k+1 ... j ... i-1
func getMaxArea(h []int) int {
	var s = new(basic.Stack)
	s.Init()

	var i int
	var maxArea int
	for i < len(h) {
		for !s.IsEmpty() && h[s.Top().(int)] >= h[i] {
			j := s.PopValue().(int)
			var k int
			if s.IsEmpty() {
				k = -1 //k+1的位置为0
			} else {
				k = s.Top().(int)
			}

			newMaxArea := (i - 1 - (k + 1) + 1) * h[j]
			if newMaxArea > maxArea {
				maxArea = newMaxArea
			}
		}
		s.Push(i)
		i++
	}

	for !s.IsEmpty() {
		//i-1的位置是len(h)-1
		j := s.PopValue().(int)
		var k int
		if s.IsEmpty() {
			k = -1 //k+1的位置为0
		} else {
			k = s.Top().(int)
		}
		newMaxArea := (len(h) - 1 - (k + 1) + 1) * h[j]
		if newMaxArea > maxArea {
			maxArea = newMaxArea
		}
	}
	fmt.Println(h, maxArea)
	return maxArea
}

func main() {
	arr := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
	}
	fmt.Println(getMaxReclFromArray(arr))
}
