package basic

import (
	"fmt"
	"testing"
)

func Test_Chast(t *testing.T) {
	ct := &Chtable{}
	ct.Init(5, HashInt, MatchInt)
	ct.Insert(1)
	ct.Insert(2)
	ct.Insert(3)
	ct.Insert(10)
	ct.Insert(10)
	fmt.Println("size", ct.size)

	fmt.Println(ct.LookUp(1))
	fmt.Println(ct.LookUp(2))
	fmt.Println(ct.LookUp(3))
	fmt.Println(ct.LookUp(10))
	ct.Remove(10)
	fmt.Println(ct.LookUp(10))
	fmt.Println(ct.size)

}
