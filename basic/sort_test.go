package basic

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25})

	t.Log("before")
	t.Log(data)
	InsertSort(data, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestBuddleSort(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25})

	t.Log("before")
	t.Log(data)
	BuddleSort(data, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestCountSort(t *testing.T) {
	data := []int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25}

	t.Log("before")
	t.Log(data)
	CountSort(data, 0, 50)
	t.Log("after")
	t.Log(data)
}

func TestRadixSort(t *testing.T) {
	data := []int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25}

	t.Log("before")
	t.Log(data)
	RadixSort(data, len(data), 2, 10)
	t.Log("after")
	t.Log(data)
}

func TestMgsort(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25})

	t.Log("before")
	t.Log(data)
	MgSort(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestQuickSort1(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25})

	t.Log("before")
	t.Log(data)
	QuickSort1(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestQuickSort2(t *testing.T) {
	data := ToInterfaceSlice([]int{1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25})

	t.Log("before")
	t.Log(data)
	QuickSort2(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestShellSort(t *testing.T) {
	a := []int{
		1, 5, 7, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 36, 25,
	}
	data := ToInterfaceSlice(a)
	l := len(data)

	ShellSort(data, l, CfInt)

	t.Log("rel ", data)
}

func TestQuick(t *testing.T) {
	arr := []int{
		1, 5, 7, 36, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 25,
	}
	QuickInt(arr)
	t.Log("rel ", arr)
}
