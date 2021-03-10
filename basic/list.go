package basic

//双向链表(可以用作单向链表),降低操作的复杂度，这里将设置好头尾指针
type List struct {
	len  int
	head *Node
	tail *Node
}

func NewListWithData(args ...interface{}) *List {
	l := new(List)
	for _, v := range args {
		l.InsertAsTail(NewNode(v))
	}
	return l
}

func (l *List) Init(node *Node) {
	if node == nil {
		return
	}
	l.len = 1
	l.head = node
	l.tail = node
}

func (l *List) Len() int {
	return l.len
}

func (l *List) IsEmpty() bool {
	return l.len == 0
}

func (l *List) Head() *Node {
	return l.head
}
func (l *List) SetHead(head *Node) {
	l.head = head
	l.head.pre = nil
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) SetTail(tail *Node) {
	l.tail = tail
	l.tail.next = nil
}

//头插法,插入节点作为链表的头节点
func (l *List) InsertAsHead(node *Node) {
	if l.len == 0 {
		l.Init(node)
		return
	}
	node.next = l.head
	l.head.pre = node
	l.head = node
	l.len++
}

//尾插法，插入节点作为链表的尾节点
func (l *List) InsertAsTail(node *Node) {
	if l.len == 0 {
		l.Init(node)
		return
	}
	node.pre = l.tail
	l.tail.next = node
	l.tail = node
	l.len++
}

//在某个节点后插入一个节点，若af为nil，表示新节点作为头节点
// 需要修改node,af.next的pre，以及af.next
//要是af是tail，需要更新tail
func (l *List) InsertAfter(af, node *Node) {
	if af == nil {
		l.InsertAsHead(node)
		return
	}
	node.next = af.next
	node.pre = af
	if af.next != nil {
		af.next.pre = node
	} else {
		l.tail = node
	}
	af.next = node
	l.len++
}

//节点前插入,若bf为nil,表示新节点作为尾节点
//需要修改node，bf原来的pre的next，以及bf的pre
//bf要是头节点，需要更新头节点
func (l *List) InsertBefore(bf, node *Node) {
	if bf == nil {
		l.InsertAsTail(node)
		return
	}
	node.next = bf
	node.pre = bf.pre
	if bf.pre != nil {
		bf.pre.next = node
	} else {
		l.head = node
	}
	bf.pre = node
	l.len++
}

//删除头节点,原来的head置为nil，head.next为新的head
func (l *List) DelHead() {
	if l.len == 1 {
		l.head, l.tail = nil, nil
		l.len = 0
		return
	}
	newHead := l.head.next
	newHead.pre = nil
	l.head = nil
	l.head = newHead
	l.len--
}

//删除尾节点，原来的tail置为nil，tail的pre为新的tail
func (l *List) DelTail() {
	if l.len == 1 {
		l.head, l.tail = nil, nil
		l.len = 0
		return
	}
	newTail := l.tail.pre
	newTail.next = nil
	l.tail = nil
	l.tail = newTail
	l.len--
}

//删除某个节点,node为链表中的某个节点,修改node.pre的next以及node.next的pre
func (l *List) DelNode(node *Node) {
	if node == l.head {
		l.DelHead()
		return
	} else if node == l.tail {
		l.DelTail()
		return
	}
	//前一个节点的后继指向node后继
	node.pre.next = node.next
	//后一个节点的前驱指向node前驱
	node.next.pre = node.pre

	node = nil
	l.len--
}

//0 for equal
//1 for big
//-1 for less
type CF func(key1, key2 interface{}) int

type Match func(key1, key2 interface{}) bool

type Hash func(key interface{}) int

func (l *List) Find(data interface{}, cf CF) *Node {
	for n := l.head; n != nil; n = n.next {
		if cf(n.data, data) == 0 {
			return n
		}
	}
	return nil
}

type CFWithNode func(n1, n2 *Node) int

func (l *List) FindWithNode(node *Node, cf CFWithNode) *Node {

	for n := l.head; n != nil; n = n.next {
		if cf(n, node) == 0 {
			return n
		}
	}
	return nil
}

//true for stop
type Invoke func(node *Node, args ...interface{}) bool

func (l *List) Traverse(iv Invoke, args ...interface{}) {
	for n := l.head; n != nil; n = n.next {
		if iv(n, args...) {
			break
		}
		if n == l.tail {
			break
		}
	}
}

//反转链表
func (l *List) Reverse() {
	//先反转头尾，尾节点作为头节点
	l.head, l.tail = l.tail, l.head

	//从原来的尾结点开始，依次修改每个节点的pre和next(调转)
	//修改完后，n跳到原来尾结点的前一个节点（既修改的下一个节点）
	for n := l.head; n != nil; n = n.next {
		n.pre, n.next = n.next, n.pre
	}
}

//作为单链表 不能用到pre指针
func (l *List) ReverseAsSingle() {
	var pre *Node
	var next *Node
	for n := l.head; n != nil; n = next {
		next = n.next
		n.next = pre
		pre = n
	}
	//此时pre为原来的tail，重设为head
	l.head, l.tail = l.tail, l.head
}

//范围反转，从from到to的发转 from从1开始
func (l *List) ReversePart(from, to int) {
	if from >= to || to > l.Len() {
		//防止panic
		return
	}

	cur := 0

	//范围的前一个节点和后一个节点
	var fPre, tPos *Node
	node1 := l.Head()

	for node1 != nil {
		cur++
		if cur == from-1 {
			fPre = node1
		}
		if cur == to+1 {
			tPos = node1
			break
		}
		node1 = node1.Next()
	}

	//node1开始反转的第一个节点，node2是node1的下一个节点
	if fPre == nil {
		node1 = l.Head()
	} else {
		node1 = fPre.Next()
	}
	node2 := node1.Next()

	//node1指向tPos
	node1.SetNext(tPos)
	//注意更新链表的尾部
	if tPos == nil {
		l.SetTail(node1)
	} else {
		tPos.SetPre(node1)
	}

	var next *Node
	//范围内反转，从node2开始，直到node2等于tPos，此时node1是反转部分的最后一个节点
	for node2 != tPos {
		next = node2.Next()
		//调整node1的pre以及node2的next
		node2.SetNext(node1)
		node1.SetPre(node2)
		node1 = node2
		node2 = next
	}

	//头节点不变，只需要调整fPre的next和node1的pre
	if fPre != nil {
		fPre.SetNext(node1)
		node1.SetPre(fPre)
		return
	}

	l.SetHead(node1)
}

//范围反转，从from到to的发转 from从1开始
func (l *List) ReversePartAsSingle(from, to int) {
	if from >= to || to > l.Len() {
		//防止panic
		return
	}

	cur := 0

	//范围的前一个节点和后一个节点
	var fPre, tPos *Node
	node1 := l.Head()

	for node1 != nil {
		cur++
		if cur == from-1 {
			fPre = node1
		}
		if cur == to+1 {
			tPos = node1
			break
		}
		node1 = node1.Next()
	}

	//node1开始反转的第一个节点，node2是node1的下一个节点
	if fPre == nil {
		node1 = l.Head()
	} else {
		node1 = fPre.Next()
	}
	node2 := node1.Next()

	//node1指向tPos
	node1.SetNext(tPos)
	//注意更新链表的尾部
	if tPos == nil {
		l.SetTail(node1)
	}

	var next *Node
	//范围内反转，从node2开始，直到node2等于tPos，此时node1是反转部分的最后一个节点
	for node2 != tPos {
		next = node2.Next()
		node2.SetNext(node1)
		node1 = node2
		node2 = next
	}

	//头节点不变，只需要调整fPre的next和node1的pre
	if fPre != nil {
		fPre.SetNext(node1)
		return
	}

	l.SetHead(node1)
}
