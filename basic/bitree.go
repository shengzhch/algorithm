package basic

import (
	"errors"
)

//单纯的二叉树
type BiTreeNode struct {
	data  interface{}
	left  *BiTreeNode
	right *BiTreeNode
}

func NewBiTreeNode(data interface{}) *BiTreeNode {
	return &BiTreeNode{
		data: data,
	}
}

type BiTree struct {
	size    int
	compare func(key1, key2 interface{}) int
	root    *BiTreeNode
}

var (
	Wrong = errors.New("Wrong")
)

func (bs *BiTree) Init(args ...interface{}) {
	bs.size = 0
	if len(args) == 1 {
		bs.compare = args[0].(func(key1, key2 interface{}) int)
	}
	return
}

//将新节点插入到二叉树中且node节点的左子结点,若node为nil且二叉树为空树，则新结点作为根结点
func (bs *BiTree) InsertAsLeftChild(node *BiTreeNode, pos *BiTreeNode) (e error) {
	if node == nil {
		if bs.size > 0 {
			return Wrong
		}

		bs.root = pos
	} else {
		if node.left != nil {
			return Wrong
		}
		node.left = pos
	}

	bs.size++
	return
}

func (bs *BiTree) InsertAsRightChild(node *BiTreeNode, pos *BiTreeNode) (e error) {
	if node == nil {
		if bs.size > 0 {
			return Wrong
		}

		bs.root = pos
	} else {
		if node.right != nil {
			return Wrong
		}
		node.right = pos
	}

	bs.size++
	return nil
}

//将指定结点node的左子树移除,node 为nil，以根结点为起始结点
func (bs *BiTree) RemoveLeftChild(node *BiTreeNode) {
	if bs.size == 0 {
		return
	}

	var pos **BiTreeNode

	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.left
	}

	//后序遍历删除所有的结点
	if *pos != nil {
		bs.RemoveLeftChild(*pos)
		bs.RemoveRightChild(*pos)
		*pos = nil
		bs.size--
	}
	return
}

func (bs *BiTree) RemoveRightChild(node *BiTreeNode) {
	if bs.size == 0 {
		return
	}

	var pos **BiTreeNode
	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.right
	}

	//后序遍历删除所有的结点
	if *pos != nil {
		bs.RemoveLeftChild(*pos)
		bs.RemoveRightChild(*pos)
		*pos = nil
		bs.size--
	}
	return
}

//将两棵二叉树合并为单棵二叉树，重新构造一个根结点
func Merge(bs, bs2 *BiTree, data interface{}) (*BiTree, error) {
	rel := &BiTree{size: 0}
	rel.Init(bs.compare)
	root := &BiTreeNode{
		data:  data,
		left:  bs.root,
		right: bs2.root,
	}
	rel.SetRoot(root)
	rel.size = 1 + bs.size + bs2.size
	bs.size = 0
	bs.root = nil
	bs2.size = 0
	bs2.root = nil
	return rel, nil
}

func (bs *BiTree) Size() int {
	return bs.size
}

func (bs *BiTree) Root() *BiTreeNode {
	return bs.root
}

func (bs *BiTree) SetRoot(root *BiTreeNode) {
	bs.root = root
}

//先序遍历 根 左 右
func (node *BiTreeNode) Preorder(l *List) {
	if !node.IsEob() {
		l.InsertAsTail(&Node{data: node.data})

		if !node.left.IsEob() {
			node.left.Preorder(l)
		}

		if !node.right.IsEob() {
			node.right.Preorder(l)
		}
	}
}

//中序遍历 左 根 右
func (node *BiTreeNode) Inorder(l *List) {
	if !node.IsEob() {
		if !node.left.IsEob() {
			node.left.Inorder(l)
		}

		l.InsertAsTail(&Node{data: node.data})

		if !node.right.IsEob() {
			node.right.Inorder(l)
		}
	}
}

//后序遍历 
func (node *BiTreeNode) Postorder(l *List) {
	if !node.IsEob() {
		if !node.left.IsEob() {
			node.left.Postorder(l)
		}

		if !node.right.IsEob() {
			node.right.Postorder(l)
		}

		l.InsertAsTail(&Node{data: node.data})
	}
}

//非递归遍历 (根结点和左孩子节点同时压栈和插入队列，走到头，然后从栈中取出一个节点，指向其右节点)
//遇到一个结点，插入到链表中，然后压栈，然后遍历左子树，左子树遍历完了，出栈，指向右结点，继续做
func (node *BiTreeNode) PreorderNo(l *List) {
	s := NewStack(nil)
	var tmp = new(BiTreeNode)

	tmp = node

	for tmp != nil || s.Size() > 0 {
		for tmp != nil {
			l.InsertAsTail(&Node{data: tmp.data})
			s.Push(tmp)
			tmp = tmp.left
		}
		if s.Size() > 0 {
			tmp = s.PopValue().(*BiTreeNode)
			tmp = tmp.right
		}
	}
}

//非递归遍历 (根结点和左孩子节点压栈，走到头，出栈一个节点，该节点插入到队列中，指针指向其右节点)
//遇到一个结点(根节点以及左子树节点)，压栈，一直找到左结点为空，出栈，插入，指向右结点，继续做
func (node *BiTreeNode) InorderNo(l *List) {
	s := NewStack(nil)
	var tmp = new(BiTreeNode)

	tmp = node

	for tmp != nil || s.Size() > 0 {
		//一路向左压栈
		for tmp != nil {
			s.Push(tmp)
			tmp = tmp.left
		}

		if s.Size() > 0 {
			tmp = s.PopValue().(*BiTreeNode)
			l.InsertAsTail(&Node{data: tmp.data})
			tmp = tmp.right
		}
	}
}

//后序遍历 左孩子 右孩子 根 ，因为根是先取到的
//两个栈来做辅助存储 依次把根结点 右儿子 左儿子压栈，出栈的顺序就是后序遍历，处理过程于前序遍历相仿
func (node *BiTreeNode) PostorderNo(l *List) {
	s1 := NewStack(nil)
	s2 := NewStack(nil)

	var tmp = new(BiTreeNode)

	tmp = node

	//根 右 左 的顺序压入s2
	for tmp != nil || s1.Size() > 0 {
		//一路向右压栈
		for tmp != nil {
			s1.Push(tmp)
			s2.Push(tmp)
			tmp = tmp.right
		}

		if s1.Size() > 0 {
			tmp = s1.PopValue().(*BiTreeNode)
			tmp = tmp.left
		}
	}

	for s2.Size() > 0 {
		tmp = s2.PopValue().(*BiTreeNode)
		l.InsertAsTail(&Node{data: tmp.data}) //插入
	}
}

func (node *BiTreeNode) PostorderNo2(l *List) {
	s1 := NewStack(nil)
	s2 := NewStack(nil)

	var tmp = node
	s1.Push(tmp)
	//一路向右压栈
	for s1.Size() > 0 {
		tmp = s1.PopValue().(*BiTreeNode)
		s2.Push(tmp)
		if !tmp.left.IsEob() {
			s1.Push(tmp.left)
		}
		if !tmp.right.IsEob() {
			s1.Push(tmp.right)
		}
	}

	for s2.Size() > 0 {
		tmp = s2.PopValue().(*BiTreeNode)
		l.InsertAsTail(&Node{data: tmp.data}) //插入
	}
}

//层级遍历 根结点入队列，然后执行循环，对头元素出队列，对头元素左儿子入队，对头元素右儿子入队
func (node *BiTreeNode) LevelorderNo(l *List) {
	q := new(Queue)
	var tmp = new(BiTreeNode)

	tmp = node

	if tmp == nil {
		return
	}
	q.EnQueue(tmp)

	for q.Size() > 0 {
		tmp = q.DeQueueWithValue().(*BiTreeNode)
		l.InsertAsTail(&Node{data: tmp.data}) //插入
		if tmp.left != nil {
			q.EnQueue(tmp.left)
		}
		if tmp.right != nil {
			q.EnQueue(tmp.right)
		}
	}
}

//判断node是不是二叉树的叶子结点
func (node *BiTreeNode) IsLeaf() bool {
	return node.left == nil && node.right == nil
}

//判断node的终止 end of bitree
func (node *BiTreeNode) IsEob() bool {
	return node == nil
}

//
func (node *BiTreeNode) Right() *BiTreeNode {
	return node.right
}

//
func (node *BiTreeNode) Left() *BiTreeNode {
	return node.left
}

//
func (node *BiTreeNode) Data() interface{} {
	return node.data
}

//
func (node *BiTreeNode) SetRight(n *BiTreeNode) {
	node.right = n
}

//
func (node *BiTreeNode) SetLeft(n *BiTreeNode) {
	node.left = n
}

//
func (node *BiTreeNode) SetData(data interface{}) {
	node.data = data
}

func (node *BiTreeNode) Show() {
}

func (node *BiTreeNode) Level() int {
	if node == nil {
		return 0
	}
	ll := node.left.Level()
	rk := node.left.Level()
	if ll > rk {
		return ll + 1
	}
	return rk + 1
}

//广度优先
func (node *BiTreeNode) BFS(l *List) {
	q := new(Queue)
	var tmp = new(BiTreeNode)
	tmp = node
	if tmp == nil {
		return
	}
	q.EnQueue(tmp)

	for q.Size() > 0 {
		n := q.DeQueueWithValue().(*BiTreeNode)
		l.InsertAsTail(NewNode(n.data))
		if n.left != nil {
			q.EnQueue(n.left)
		}
		if n.right != nil {
			q.EnQueue(n.right)
		}
	}
}

//深度优先 递归写法
func (node *BiTreeNode) DFS(l *List) {
	var tmp = new(BiTreeNode)
	tmp = node
	if tmp != nil {
		l.InsertAsTail(NewNode(tmp.data))
		if tmp.left != nil {
			tmp.left.DFS(l)
		}
		if tmp.right != nil {
			tmp.right.DFS(l)
		}
	}
}

//深度优先 实现是先序遍历
func (node *BiTreeNode) DFSCircle(l *List) {
	var tmp = new(BiTreeNode)
	tmp = node
	if tmp != nil {
		l.InsertAsTail(NewNode(tmp.data))
		if tmp.left != nil {
			tmp.left.DFS(l)
		}
		if tmp.right != nil {
			tmp.right.DFS(l)
		}
	}
}

//广度优先方式插入
func (tree *BiTree) BFSInsert(v interface{}) {
	if tree.root == nil {
		tree.root = &BiTreeNode{data: v}
		return
	}

	//
	q := make([]*BiTreeNode, 0)
	q = append(q, tree.root)
	for len(q) > 0 {
		//队列
		curNode := q[0] // 非递归分层遍历
		q = q[1:]       // pop

		if curNode.left != nil { // 左子节点
			q = append(q, curNode.left)
		} else {
			curNode.left = &BiTreeNode{data: v} // 找到合适的插入地点
			return
		}

		if curNode.right != nil {
			q = append(q, curNode.right)
		} else {
			curNode.right = &BiTreeNode{data: v}
			return
		}
	}
}
