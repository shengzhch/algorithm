package basic

<<<<<<< HEAD
=======
import "fmt"

//链表中结点，这里包含两个指针，指向前驱和后继
>>>>>>> origin/master
type Node struct {
	data interface{}
	pre  *Node
	next *Node
}

<<<<<<< HEAD
=======
func NewNode(d interface{}) *Node {
	return &Node{data: d}
}

>>>>>>> origin/master
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
<<<<<<< HEAD
=======

func (n *Node) PrintAsSingleList() {
	cur := n
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}

	return
}
>>>>>>> origin/master
