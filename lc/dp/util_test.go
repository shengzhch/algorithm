package main

import "testing"

func Test_MaxDG(t *testing.T) {
	a := []int{1, 3, 4, 2, 5}
	t.Log(len(a), cap(a))
	t.Log(maxDG2(a))
	t.Log(a)
	t.Log(len(a), cap(a))
}
