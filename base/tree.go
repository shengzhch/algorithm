package base

type Tree interface {
	Put(k, v int)  //新增或修改
	Get(k int) int //查询
	Delete(k int)  //删除
	Size() int     //树的大小
	Min() int      //最小键
	DeleteMin()    //删除最小键
}

type node struct {
	val         ValueInterface
	num         int
	left, right *node //左右子节点
}

//创建节点
func newNode(v ValueInterface) *node {
	return &node{val: v, num: 1}
}

//获取以nd为根节点的树的大小
func size(nd *node) int {
	if nd == nil {
		return 0
	}
	return nd.num
}

//更新以nd为根节点的树的大小
func updateSize(nd *node) {
	if nd == nil {
		return
	}
	nd.num = size(nd.left) + size(nd.right) + 1
}

func put(nd *node, v ValueInterface) (*node, bool) {
	if nd == nil {
		return newNode(v), true
	}
	hasNew := false //检测是否创建了新节点

	if v.Less(nd.val) {
		nd.left, hasNew = put(nd.left, v)
	} else {
		nd.right, hasNew = put(nd.right, v)
	}

	if hasNew { //如果创建了新节点就更新树的大小
		updateSize(nd)
	}
	return nd, hasNew
}
