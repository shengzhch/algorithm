package basic

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	s := NewQueue()
	s.EnQueue(1)
	s.EnQueue(2)
	s.EnQueue(3)
	//1 2 3
	s.Traverse(PrintNode)
	fmt.Println(s.Size())             //3
	fmt.Println(s.DeQueueWithValue()) //1
	fmt.Println(s.DeQueueWithValue()) //2
	fmt.Println(s.Size())             //1
	fmt.Println(s.DeQueueWithValue()) //3
	s.EnQueue(4)
	fmt.Println(s.Size()) //1
}
