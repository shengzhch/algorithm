package basic

//循环链表
type CList List

func NewClist(node *Node) CList {
	l := new(List)
	l.Init(node)
	l.tail.next = l.head
	return CList(*l)
}

//在ele后插入一个新的元素
func (l *CList) InsertAfter(ele *Node, newNode *Node) {
	if l.len == 0 {
		newNode.next = newNode
		l.head = newNode
		l.tail = newNode
	} else {
		ele.next.pre = newNode
		newNode.next = ele.next
		newNode.pre = ele
		ele.next = newNode
	}

	//重新设置tail
	if ele == l.tail {
		l.tail = newNode
	}

	l.len++
	return
}

//在ele后删除一个新的元素
func (l *CList) DeleteAfter(ele *Node) {
	if l.len == 0 {
		return
	}

	//单个节点
	if ele.next == ele {
		l.head = nil
		l.tail = nil
	} else {
		next := ele.next
		next.next.pre = ele
		ele.next = next.next

		//删除节点是头结点，头节点后移
		if next == l.head {
			l.head = next.next
		}
		//删除节点是尾节点，尾节点前移
		if next == l.tail {
			l.tail = ele
		}
		next = nil
	}
	l.len--
	return
}

//头节点
func (s *CList) Head() *Node {
	return s.head
}

//尾节点
func (s *CList) Tail() *Node {
	return s.tail
}

//遍历
func (s *CList) Traverse(iv Invoke, args ...interface{}) {
	(*List)(s).Traverse(iv, args...)
}
