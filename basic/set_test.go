package basic

import (
	"fmt"
	"testing"
)

func Test_Set(t *testing.T) {
	s := NewSet(nil)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(1)

	fmt.Println(s.Belong(1)) //true
	fmt.Println(s.Belong(2)) //true
	fmt.Println(s.Belong(4)) //false

	fmt.Println(s.Size())//3

	s2 := NewSet(nil)

	s2.Insert(2)
	s2.Insert(4)

	fmt.Println(s.Contain(&s2)) //false

	fmt.Println("并集")
	s.Union(&s2).Traverse(PrintNode) //1234
	fmt.Println("s-s2差集")
	s.Difference(&s2).Traverse(PrintNode) //13
	fmt.Println("s2-s1差集")
	s2.Difference(&s).Traverse(PrintNode) //4
}
