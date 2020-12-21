package basic

//栈 后进先出。栈顶进行插入删除。 用链表表示，表头表示栈顶
type Stack List

func NewStack(node *Node) Stack {
	l := new(List)
	l.Init(node)
	return Stack(*l)
}

func (s *Stack) Push(data interface{}) {
	(*List)(s).InsertAsHead(&Node{data: data})
}

func (s *Stack) Pop() {
	(*List)(s).DelHead()
}

func (s *Stack) PopValue() interface{} {
	rel := s.head.Data()
	(*List)(s).DelHead()
	return rel
}

func (s *Stack) Top() interface{} {
	return s.head.Data()
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}
