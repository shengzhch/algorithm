package basic

import (
	"math/rand"
	"testing"
)

func TestInsertSort(t *testing.T) {
	var data = make([]interface{}, 0, 10)

	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(10))
	}

	t.Log("before")
	t.Log(data)
	InsertSort(data, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestBuddleSort(t *testing.T) {
	var data = make([]interface{}, 0, 10)

	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(10))
	}

	t.Log("before")
	t.Log(data)
	BuddleSort(data, CfInt)
	t.Log("after")
	t.Log(data)
}

func TestCountSort(t *testing.T) {
	var data = make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		data = append(data, i)
	}

	t.Log("before")
	t.Log(data)
	CountSort(data, 0, 10)
	t.Log("after")
	t.Log(data)
}

func TestRadixSort(t *testing.T) {
	var data = make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(1000))
	}

	t.Log("before")
	t.Log(data)
	RadixSort(data, 10, 3, 10)
	t.Log("after")
	t.Log(data)
}

func TestMgsort(t *testing.T) {
	var data = make([]interface{}, 0, 10)

	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(100))
	}

	t.Log("before")
	t.Log(data)
	MgSort(data, 0, len(data)-1, CfInt)
	t.Log("after")
	t.Log(data)
}
