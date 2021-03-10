package basic

//队列 先进先出
//对头删除,对尾插入
//链表实现，头节点进行出队，尾节点进行入队
type Queue List

func NewQueue() Queue {
	l := new(List)
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

//是否为空
func (s *Queue) IsEmpty() bool {
	return s.len == 0
}

//遍历
func (s *Queue) Traverse(iv Invoke, args ...interface{}) {
	(*List)(s).Traverse(iv, args...)
}
