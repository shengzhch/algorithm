package basic

import (
	"fmt"
	"testing"
)

func Test_Heap(t *testing.T) {
	h := new(Heap)
	h.Init(CfInt)

	h.Insert(1)
	h.Insert(4)
	h.Insert(2)
	h.Insert(3)
	fmt.Printf("top %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
	h.Insert(5)
	fmt.Printf("insert 5 after\ntop %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
	h.Extract()
	fmt.Printf("Extract after\ntop %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
	h.Extract()
	fmt.Printf("Extract again after\ntop %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
	h.Extract()
	fmt.Printf("Extract again after\ntop %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
	h.Extract()
	fmt.Printf("Extract again after\ntop %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
	h.Extract()
	fmt.Printf("Extract again after\ntop %v\n", h.Top())
	fmt.Printf("size %v\n", h.Size())
	fmt.Printf("data %v\n", h.Data())
}
