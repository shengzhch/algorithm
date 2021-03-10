package basic

import (
	"fmt"
	"testing"
)

func Test_Ohash(t *testing.T) {
	ot := &Ohtable{}
	ot.Init(5, HashInt, HashInt, MatchInt)
	ot.Insert(1)
	ot.Insert(2)
	ot.Insert(3)
	ot.Insert(10)
	ot.Insert(10)
	fmt.Println("size", ot.size)

	fmt.Println(ot.LookUp(1))
	fmt.Println(ot.LookUp(2))
	fmt.Println(ot.LookUp(3))
	fmt.Println(ot.LookUp(4))
	fmt.Println(ot.LookUp(10))
	ot.Remove(10)
	fmt.Println(ot.LookUp(10))
	fmt.Println(ot.size)

}
