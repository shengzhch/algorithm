package base

import "sync"

//栈 ：栈顶插入删除。后进先出。由于有切片存储数据，故采用0表示栈底元素下标，top为栈顶元素下标
//同样可以用链表的方式来表示栈，这里不再给出定义

type Stack struct {
	values []ValueInterface
	top    int
	lock   *sync.RWMutex
	cap    int
}

func NewStack(cap int) *Stack {
	s := &Stack{}
	s.values = make([]ValueInterface, 0)
	s.top = -1
	s.cap = cap
	s.lock = new(sync.RWMutex)
	return s
}



func (s *Stack) Top() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.top
}

func (s *Stack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.top == -1
}

func (s *Stack) IsFull() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.cap < 0 {
		return false
	}
	return (s.top + 1) == s.cap
}

func (s *Stack) Pop() (e ValueInterface) {
	s.lock.Lock()
	defer s.lock.Unlock()

	e, s.values = s.values[s.top], s.values[:s.top]
	s.top--
	return
}

func (s *Stack) Push(e ValueInterface) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.values = append(s.values, e)
	s.top++
}

func (s *Stack) TopValue() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.values[s.top]
}
