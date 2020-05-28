package basic

//队列 先进先出 对头删除,对尾插入 单链表表示队列，头节点进行出队操作，尾节点进行入队操作
type Queue List

func NewQueue(node *Node) Queue {
	l := new(List)
	l.Init(node)
	return Queue(*l)
}

//入队
func (s *Queue) EnQueue(data interface{}) {
	(*List)(s).InsertAsTail(&Node{data: data})
}

//出队
func (s *Queue) DeQueue() {
	(*List)(s).DelHead()
}

//
func (s *Queue) DeQueueWithValue() interface{} {
	rel := List(*s).head
	(*List)(s).DelHead()
	return rel.Data()
}

func (s *Queue) Size() int {
	return s.len
}

func (s *Queue) Head() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.Data()
}

func (s *Queue) Tail() interface{} {
	if s.tail == nil {
		return nil
	}
	return s.tail.Data()
}
