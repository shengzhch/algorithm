package base

import (
	"errors"
	"fmt"
)

//链表(双向链表)
type LinkList struct {
	len  int
	head *Node
	rear *Node
}

func NewLinkList() *LinkList {
	return &LinkList{len: 0}
}

func (l *LinkList) String() string {
	var rel string
	l.Map(func(n *Node) {
		rel += fmt.Sprintf("-> %v ", n.value)
	})
	return rel
}

func (l *LinkList) Len() int {
	return l.len
}

func (l *LinkList) IsEmpty() bool {
	return l.len == 0
}

//获得链表第index节点 head 为第1个节点
func (l *LinkList) Get(index int) (*Node, error) {
	if index > l.Len() {
		return nil, errors.New("Index out of range")
	}
	node := l.head
	for i := 1; i < index; i++ {
		node = node.next
	}
	return node, nil
}

func (l *LinkList) Clear() {
	l.len = 0
	l.head = nil
	l.rear = nil
}

//头插
func (l *LinkList) HeadInsert(v ValueInterface) {
	node := NewNode(v)

	if l.Len() == 0 {
		l.head = node
		l.rear = l.head
	} else {
		curHead := l.head
		curHead.pre = node
		node.next = curHead
		l.head = node
	}

	l.len++
}

//尾插
func (l *LinkList) RearInsert(value ValueInterface) {
	node := NewNode(value)

	if l.Len() == 0 {
		l.head = node
		l.rear = l.head
	} else {
		curRear := l.rear
		curRear.next = node
		node.pre = curRear
		l.rear = node
	}

	l.len++
}

//指定位置插入
func (l *LinkList) AddByIndex(value ValueInterface, index int) error {
	if index > l.Len() {
		return errors.New("Index out of range")
	}

	node := NewNode(value)

	if index == 0 {
		l.HeadInsert(value)
		return nil
	}

	if index == l.Len()+1 {
		l.RearInsert(value)
		return nil
	}

	nextNode, _ := l.Get(index)
	prevNode := nextNode.pre

	prevNode.next = node
	node.pre = prevNode

	nextNode.pre = node
	node.next = nextNode

	l.len++

	return nil
}

//根据value值移除移除第一个为该值得节点
func (l *LinkList) Remove(value ValueInterface) error {
	if l.Len() == 0 {
		return errors.New("Empty LinkList")
	}

	if l.head.value.Equal(value) {
		l.head = l.head.next
		l.head.pre = nil
		l.len--
		return nil
	}

	for n := l.head.next; n != nil; n = n.next {
		if n.value.Equal(value) {
			n.pre.next = n.next
			n.next.pre = n.pre
			l.len--
			return nil
		}
	}
	return errors.New("Node not found")
}

//移除第index个元素
func (l *LinkList) RemoveByIndex(index int) error {
	if index > l.Len() || index < 1 {
		return errors.New("Index out of range")
	}
	if index == 1 {
		l.head = l.head.next
		l.head.pre = nil
		l.len--
		return nil
	}
	if index == l.len {
		l.rear = l.rear.pre
		l.rear.next = nil
		l.len--
		return nil
	}

	curNode, _ := l.Get(index)
	curNode.pre.next = curNode.next
	curNode.next.pre = curNode.pre
	l.len--

	return nil
}

//根据value确定为链表中的第几个元素
func (l *LinkList) Find(node *Node) (int, error) {
	if l.Len() == 0 {
		return 0, errors.New("Empty LinkList")
	}

	index := 0
	found := -1
	l.Map(func(n *Node) {
		index++
		if node.value.Equal(node.value) {
			found = index
		}
	})

	if found == -1 {
		return 0, errors.New("Item not found")
	}

	return found, nil
}

//链接两个链表
func (l *LinkList) Concat(k *LinkList) {
	l.rear.next, k.head.pre = k.head, l.rear
	l.rear = k.rear
	l.len += k.len
}

//双向链表链表反转
func (l *LinkList) Reverse() {
	if l.len <= 1 {
		return
	}
	rear := l.head
	cur := l.head
	curNext := cur.next

	for i := 1; i < l.len; i++ {
		t := curNext.next
		curNext.next = cur
		cur.pre = curNext

		cur = curNext
		curNext = t
	}

	l.head = cur
	l.head.pre = nil
	l.rear = rear
	l.rear.next = nil
	return
}

//遍历
func (LinkList *LinkList) Map(f func(node *Node)) {
	for node := LinkList.head; node != nil; node = node.next {
		f(node)
	}
}
