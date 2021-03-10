package basic

//循环链表
type CList List

func NewClist() CList {
	l := new(List)
	return CList(*l)
}

func (l *CList) InsertAsTail(newNode *Node) {
	if l.len == 0 {
		newNode.next = newNode
		newNode.next = newNode
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.pre = l.tail
		newNode.next = l.tail.next
		l.tail.next.pre = newNode
		l.tail.next = newNode
		l.tail = newNode
	}

	l.len++
	return
}

// 删除
func (l *CList) Delete(ele *Node) {
	if l.len == 0 {
		return
	}

	//单个节点
	if ele.next == ele {
		l.head = nil
		l.tail = nil
	} else {
		ele.next.pre = ele.pre
		ele.pre.next = ele.next

		//删除节点是头结点，头节点后移
		if ele == l.head {
			l.head = ele.next
		}
		//删除节点是尾节点，尾节点前移
		if ele == l.tail {
			l.tail = ele.pre
		}
	}
	ele = nil
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
