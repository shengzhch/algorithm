package basic

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 6, 4, 3, 5, 2})

	t.Log("before")
	t.Log(data)
	InsertSort(data, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestBuddleSort(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 6, 4, 3, 5, 2})

	t.Log("before")
	t.Log(data)
	BuddleSort(data, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestCountSort(t *testing.T) {
	data := []int{1, 6, 4, 3, 5, 2}

	t.Log("before")
	t.Log(data)
	CountSort(data, 0, 10)
	t.Log("after")
	t.Log(data)
}

func TestRadixSort(t *testing.T) {
	data := []int{1, 6, 4, 18, 3, 5, 2, 21, 2}

	t.Log("before")
	t.Log(data)
	RadixSort(data, len(data), 2, 10)
	t.Log("after")
	t.Log(data)
}

func TestMgsort(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 6, 4, 3, 5, 2})

	t.Log("before")
	t.Log(data)
	MgSort(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestQuickSort1(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 6, 4, 3, 5, 2})

	t.Log("before")
	t.Log(data)
	QuickSort1(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestQuickSort2(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 6, 4, 3, 5, 2})

	t.Log("before")
	t.Log(data)
	QuickSort2(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}
