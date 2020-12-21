package basic

type Node struct {
	data interface{}
	pre  *Node
	next *Node
}

func (n *Node) SetPre(pre *Node) {
	n.pre = pre
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) Pre() *Node {
	return n.pre
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Data() interface{} {
	return n.data
}
