package basic

//栈 后进先出。
//栈顶进行插入删除。
//用链表实现，head标示栈顶，进行入栈和出栈
type Stack List

func NewStack(node *Node) Stack {
	l := new(List)
	l.Init(node)
	return Stack(*l)
}

//入栈 头插法
func (s *Stack) Init() {
	(*List)(s).Init(nil)
}

//入栈 头插法
func (s *Stack) Push(data interface{}) {
	(*List)(s).InsertAsHead(&Node{data: data})
}

//出栈 删除头节点
func (s *Stack) Pop() {
	(*List)(s).DelHead()
}

//出栈并且返回值
func (s *Stack) PopValue() interface{} {
	rel := s.head.Data()
	(*List)(s).DelHead()
	return rel
}

//返回栈顶
func (s *Stack) Top() interface{} {
	return s.head.Data()
}

//返回大小
func (s *Stack) Size() int {
	return s.len
}

//是否为空
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

//遍历
func (s *Stack) Traverse(iv Invoke, args ...interface{}) {
	(*List)(s).Traverse(iv, args...)
}
