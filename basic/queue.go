package basic

//队列 先进先出
//对头删除,对尾插入
//链表实现，头节点进行出队，尾节点进行入队
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

//出队并且返回值
func (s *Queue) DeQueueWithValue() interface{} {
	rel := List(*s).head
	(*List)(s).DelHead()
	return rel.Data()
}

//返回大小
func (s *Queue) Size() int {
	return s.len
}

//返回对头
func (s *Queue) Head() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.Data()
}

//返回队尾
func (s *Queue) Tail() interface{} {
	if s.tail == nil {
		return nil
	}
	return s.tail.Data()
}

//是否为空
func (s *Queue) IsEmpty() bool {
	return s.len == 0
}

//遍历
func (s *Queue) Traverse(iv Invoke, args ...interface{}) {
	(*List)(s).Traverse(iv, args...)
}
