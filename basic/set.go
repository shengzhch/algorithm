package basic

//集合
type Set List

func NewSet(node *Node) Set {
	l := new(List)
	l.Init(node)
	return Set(*l)
}

//属于
func (s *Set) Belong(data interface{}) bool {
	for m := s.head; m != nil; m = m.next {
		if m.data == data {
			return true
		}
	}
	return false
}

//插入
func (s *Set) Insert(data interface{}) {
	if !s.Belong(data) {
		(*List)(s).InsertAsTail(&Node{data: data})
	}
}

//删除
func (s *Set) Remove(data interface{}) {
	var find *Node
	for m := s.head; m != nil; m = m.next {
		if m.data == data {
			find = m
			break
		}
	}
	if find == nil {
		return
	}
	(*List)(s).DelNode(find)
	return
}

//并集
func (s *Set) Union(s2 *Set) *Set {
	var rel = new(Set)
	for m := s.head; m != nil; m = m.next {
		(*List)(rel).InsertAsTail(&Node{data: m.data})
	}
	for m := s2.head; m != nil; m = m.next {
		if rel.Belong(m.data) {
			continue
		} else {
			(*List)(rel).InsertAsTail(&Node{data: m.data})
		}
	}
	return rel
}

//交集
func (s *Set) Intersection(s2 *Set) *Set {
	var rel = new(Set)
	for m := s.head; m != nil; m = m.next {
		if s2.Belong(m.data) {
			(*List)(rel).InsertAsTail(&Node{data: m.data})
		}
	}
	return rel

}

//差集 s - s2
func (s *Set) Difference(s2 *Set) *Set {
	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		if !s2.Belong(m.data) {
			(*List)(rel).InsertAsTail(&Node{data: m.data})
		}
	}
	return rel
}

//包含
func (s *Set) Contain(s2 *Set) bool {
	if s2.Size() > s.Size() {
		return false
	}
	for m := s2.head; m != nil; m = m.next {
		if !s.Belong(m.data) {
			return false
		}
	}
	return true
}

//相等
func (s *Set) Equal(s2 *Set) bool {
	if s.Size() != s2.Size() {
		return false
	}
	for m := s2.head; m != nil; m = m.next {
		if !s.Belong(m.data) {
			return false
		}
	}
	return true
}

func (s *Set) Size() int {
	return (*List)(s).len
}

func (s *Set) Traverse(iv Invoke, args ...interface{}) {
	(*List)(s).Traverse(iv, args...)
}
