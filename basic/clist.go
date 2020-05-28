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
func (l *CList) InsertAfter(ele *Node, value interface{}) {
	node := &Node{
		data: value,
	}

	if l.len == 0 {
		node.next = node
		l.head = node
		l.tail = node
	} else {
		node.next = ele.next
		ele.next = node
	}

	//重新设置tail
	if ele == l.tail {
		l.tail = node
	}

	l.len++
	return
}

//删除元素ele后的元素
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
		ele.next = ele.next.next
		//删除节点是头结点，头节点后移
		if next == l.head {
			l.head = next.next
		}
		//删除节点是尾节点，尾节点前移
		if next == l.tail {
			l.tail = ele
		}
	}
	l.len--
	return
}
