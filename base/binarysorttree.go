package base

//二叉排序树
type BinarySortTree struct {
	root *node
}

//创建二叉查找树
func NewBinaryTree() *BinarySortTree {
	return &BinarySortTree{}
}

//Insert
func (bt *BinarySortTree) Insert(v ValueInterface) {
	bt.root, _ = put(bt.root, v)
}
