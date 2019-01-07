package base

import (
	"fmt"
	"errors"
)

//链表(单向链表) 只有有头结点
type SinglyLinkList struct {
	len  int
	head *Node
}

func NewSinglyLinkList() *SinglyLinkList {
	return &SinglyLinkList{len: 0}
}

func (l *SinglyLinkList) String() string {
	var rel string
	l.Map(func(n *Node) {
		rel += fmt.Sprintf("-> %v ", n.value)
	})
	return rel
}

func (l *SinglyLinkList) Len() int {
	return l.len
}

func (l *SinglyLinkList) IsEmpty() bool {
	return l.len == 0
}

//获得链表第index节点
func (l *SinglyLinkList) Get(index int) (*Node, error) {
	if index > l.Len() {
		return nil, errors.New("Index out of range")
	}
	node := l.head
	for i := 1; i < index; i++ {
		node = node.next
	}
	return node, nil
}

func (l *SinglyLinkList) Clear() {
	l.len = 0
	l.head = nil
}

//头插
func (l *SinglyLinkList) HeadInsert(v ValueInterface) {
	node := NewNode(v)

	if l.Len() == 0 {
		l.head = node
	} else {
		curHead := l.head
		node.next = curHead
		l.head = node
	}

	l.len++
}

//尾插
func (l *SinglyLinkList) RearInsert(value ValueInterface) {
	node := NewNode(value)

	if l.Len() == 0 {
		l.head = node
	} else {
		curRear, _ := l.Get(l.len)
		curRear.next = node
	}

	l.len++
}

//指定位置插入
func (l *SinglyLinkList) Add(value ValueInterface, index int) error {
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

	preNode, _ := l.Get(index - 1)
	node.next = preNode.next
	preNode.next = node
	l.len++

	return nil
}

//根据value值移除移除第一个为该值得节点
func (l *SinglyLinkList) Remove(value ValueInterface) error {
	if l.Len() == 0 {
		return errors.New("Empty SinglyLinkList")
	}

	if l.head.value.Equal(value) {
		l.head = l.head.next
		l.len--
		return nil
	}

	for n := l.head; n.next != nil; n = n.next {
		if n.next.value.Equal(value) {
			n.next = n.next.next
			l.len--
			return nil
		}
	}
	return errors.New("Node not found")
}

//移除第index个元素
func (l *SinglyLinkList) RemoveByIndex(index int) error {
	if index > l.Len() || index < 1 {
		return errors.New("Index out of range")
	}
	if index == 1 {
		l.head = l.head.next
		l.len--
		return nil
	}

	curNode, _ := l.Get(index - 1)
	curNode.next = curNode.next.next
	l.len--

	return nil
}

//根据value确定为链表中的第几个元素
func (l *SinglyLinkList) Find(node *Node) (int, error) {
	if l.Len() == 0 {
		return 0, errors.New("Empty SinglyLinkList")
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
func (l *SinglyLinkList) Concat(k *SinglyLinkList) {
	curRear, _ := l.Get(l.len)
	curRear.next = k.head
	l.len += k.len
}

//链表反转
func (l *SinglyLinkList) Reverse() {
	if l.len <= 1 {
		return
	}
	cur := l.head
	curNext := l.head.next
	for i := 1; i < l.len; i++ {
		t := curNext.next
		curNext.next = cur
		cur = curNext
		curNext = t
	}
	l.head = cur
}

//遍历
func (SinglyLinkList *SinglyLinkList) Map(f func(node *Node)) {
	for node := SinglyLinkList.head; node != nil; node = node.next {
		f(node)
	}
}
