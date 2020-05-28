package basic

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	s := NewQueue(&Node{data: 1})
	s.EnQueue(2)
	s.EnQueue(3)
	fmt.Println(s.Size()) //3
	fmt.Println(s.DeQueueWithValue()) //1
	fmt.Println(s.DeQueueWithValue()) //2
	fmt.Println(s.Size())             //1
	fmt.Println(s.Head())             //3
	fmt.Println(s.DeQueueWithValue()) //3
	fmt.Println(s.Head())             //nil
	s.EnQueue(4)
	fmt.Println(s.Head()) //4
	fmt.Println(s.Tail()) //4
	fmt.Println(s.Size()) //1
}
