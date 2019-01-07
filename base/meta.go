package base

type ValueInterface interface {
	Equal(than ValueInterface) bool
	Less(than ValueInterface) bool
	Big(than ValueInterface) bool
	LessEqual(than ValueInterface) bool
	BigEqual(than ValueInterface) bool
	Add(ValueInterface) ValueInterface
	Sub(ValueInterface) ValueInterface
}

type Node struct {
	value ValueInterface
	pre   *Node
	next  *Node
}

func NewNode(v ValueInterface) *Node {
	return &Node{value: v}
}

var Sum = func(objs ... ValueInterface) ValueInterface {
	var rel ValueInterface
	for i, obj := range objs {
		if i == 0 {
			rel = obj
		} else {
			rel = rel.Add(obj)
		}
	}
	return rel
}

var Max = func(objs ... ValueInterface) ValueInterface {
	var rel ValueInterface
	for i, obj := range objs {
		if i == 0 {
			rel = obj
		} else {
			if obj.Big(rel) {
				rel = obj
			}
		}
	}
	return rel
}

var Min = func(objs ... ValueInterface) ValueInterface {
	var rel ValueInterface
	for i, obj := range objs {
		if i == 0 {
			rel = obj
		} else {
			if obj.Less(rel) {
				rel = obj
			}
		}
	}
	return rel
}

type IntValue int

func (i IntValue) Equal(than ValueInterface) bool {
	return i == than.(IntValue)
}

func (i IntValue) Less(than ValueInterface) bool {
	return i < than.(IntValue)
}

func (i IntValue) Add(than ValueInterface) ValueInterface {
	return i + than.(IntValue)
}

func (i IntValue) Sub(than ValueInterface) ValueInterface {
	return i - than.(IntValue)
}

func (i IntValue) Big(than ValueInterface) bool {
	return i > than.(IntValue)
}

func (i IntValue) LessEqual(than ValueInterface) bool {
	return i <= than.(IntValue)
}

func (i IntValue) BigEqual(than ValueInterface) bool {
	return i >= than.(IntValue)
}
