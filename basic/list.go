package basic

//双向链表
type List struct {
	len  int
	head *Node
	tail *Node
}

func (l *List) Init(node *Node) {
	if node == nil {
		return
	}
	l.len = 1
	l.head = node
	l.tail = node

}

func (l *List) Len() int {
	return l.len
}

func (l *List) Head() *Node {
	return l.head
}
func (l *List) SetHead(head *Node) {
	l.head = head
	l.head.pre = nil
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) SetTail(tail *Node) {
	l.tail = tail
	l.tail.next = nil
}

//头插法
func (l *List) InsertAsHead(node *Node) {
	if l.len == 0 {
		l.Init(node)
		return
	}
	node.next = l.head
	l.head.pre = node
	l.head = node
	l.len++
	if l.len == 1 {
		l.tail = l.head
	}
}

//尾插法
func (l *List) InsertAsTail(node *Node) {
	if l.len == 0 {
		l.Init(node)
		return
	}
	node.pre = l.tail
	l.tail.next = node
	l.tail = node
	l.len++
}

//节点后插入
func (l *List) InsertAfter(af, node *Node) {
	node.next = af.next
	node.pre = node
	af.next.pre = node
	af.next = node
	l.len++
}

//节点前插入
func (l *List) InsertBefore(bf, node *Node) {
	node.next = bf
	node.pre = bf.pre
	bf.pre.next = node
	bf.pre = node

	if bf == l.head {
		l.head = node
	}
	l.len++
}

//删除头
func (l *List) DelHead() {
	if l.len == 1 {
		l.head, l.tail = nil, nil
		l.len = 0
		return
	}
	l.head = l.head.next
	l.head.pre = nil
	l.len--
}

//删除尾
func (l *List) DelTail() {
	if l.len == 1 {
		l.head, l.tail = nil, nil
		l.len = 0
		return
	}
	l.tail = l.tail.pre
	l.tail.next = nil
	l.len--

}

//删除某个节点
func (l *List) DelNode(node *Node) {
	if node == l.head {
		l.DelHead()
		return
	} else if node == l.tail {
		l.DelTail()
		return
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	node = nil
	l.len--
}

//0 equal 1 big -1 less
type CF func(key1, key2 interface{}) int

func (l *List) Find(data interface{}, cf CF) *Node {

	for n := l.head; n != nil; n = n.next {
		if cf(n.data, data) == 0 {
			return n
		}
	}

	return nil
}

//true for stop
type Invoke func(node *Node, args ...interface{}) bool

func (l *List) Traverse(iv Invoke, args ...interface{}) {
	for n := l.head; n != nil; n = n.next {
		if iv(n, args...) {
			break
		}
		if n == l.tail {
			break
		}
	}
}

func (l *List) Reverse() {

	//先反转头尾
	l.head, l.tail = l.tail, l.head

	for n := l.head; n != nil; n = n.next {
		n.pre, n.next = n.next, n.pre
	}
}

//作为单链表 不能用到pre指针

func (l *List) ReverseAsSingle() {
	var pre *Node
	var next *Node
	l.tail = l.head
	for n := l.head; n != nil; {
		next = n.next
		n.next = pre
		pre = n
		n = next
	}
	l.head = pre
}
