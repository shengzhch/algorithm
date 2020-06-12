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
