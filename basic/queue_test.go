package basic

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
<<<<<<< HEAD
	s := NewQueue(&Node{data: 1})
=======
	s := NewQueue()
	s.EnQueue(1)
>>>>>>> origin/master
	s.EnQueue(2)
	s.EnQueue(3)
	//1 2 3
	s.Traverse(PrintNode)
	fmt.Println(s.Size())             //3
	fmt.Println(s.DeQueueWithValue()) //1
	fmt.Println(s.DeQueueWithValue()) //2
	fmt.Println(s.Size())             //1
<<<<<<< HEAD
	fmt.Println(s.Head())             //3
	fmt.Println(s.DeQueueWithValue()) //3
	fmt.Println(s.Head())             //nil
	s.EnQueue(4)
	fmt.Println(s.Head()) //4
	fmt.Println(s.Tail()) //4
=======
	fmt.Println(s.DeQueueWithValue()) //3
	s.EnQueue(4)
>>>>>>> origin/master
	fmt.Println(s.Size()) //1
}
