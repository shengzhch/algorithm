/*
一棵深度为k且有 2^k-1 个结点的二叉树称为满二叉树

如果对满二叉树的结点进行编号, 约定编号从根结点起, 自上而下, 自左而右。
则深度为k的, 有n个结点的二叉树, 当且仅当其每一个结点都与深度为k的满二叉树中编号从1至n的结点一一对应时, 称之为完全二叉树。

满二叉树是完全二叉树的特殊形态, 即如果一棵二叉树是满二叉树, 则它必定是完全二叉树

堆
子节点的值比父结点的值小，根结点是树中最大的结点，称为最大值堆。
堆
用数组的方式存储数据，二叉树的层级遍历，左平衡树
<<<<<<< HEAD
=======

本堆的实现是逻辑意义上的最大堆实现
>>>>>>> origin/master
*/

package basic

type Heap struct {
	size    int
	compare func(key1, key2 interface{}) int
	tree    []interface{}
}

func (h *Heap) Init(cf CF) {
	h.size = 0
	h.tree = make([]interface{}, 0, 0)
	h.compare = cf
	return
}

<<<<<<< HEAD
var (
=======
//下标从0开始
var (
	//父节点下标
>>>>>>> origin/master
	heapParent = func(npos int) int {
		return (npos - 1) / 2
	}

<<<<<<< HEAD
=======
	//左孩子
>>>>>>> origin/master
	heapLeft = func(npos int) int {
		return (npos * 2) + 1
	}

<<<<<<< HEAD
=======
	//右孩子
>>>>>>> origin/master
	heapRight = func(npos int) int {
		return (npos * 2) + 2
	}
)

//插入 插入后保证堆的结构
func (h *Heap) Insert(data interface{}) {
	var ipos, ppos int

	h.tree = append(h.tree, data)

	ipos = h.size
<<<<<<< HEAD
=======
	//当前节点的父节点
>>>>>>> origin/master
	ppos = heapParent(ipos)

	//一直查找父节点，不满足堆的性质，就交换位置
	for ipos > 0 && h.compare(h.tree[ppos], h.tree[ipos]) < 0 {
		h.tree[ppos], h.tree[ipos] = h.tree[ipos], h.tree[ppos]
		ipos = ppos
		ppos = heapParent(ipos)
	}
	h.size++
}

//提取堆顶元素，取出后调整数据，是其满足堆的性质
func (h *Heap) Extract() interface{} {
	if h.size == 0 {
		return nil
	}

	data := h.tree[0]

	//取出最后一个元素
	save := h.tree[h.size-1]

	if h.size > 1 {
		//调整数组
		h.tree = h.tree[:h.size-1]
		h.size--
	} else {
		h.tree = nil
		h.size = 0
		return data
	}

	//最后一个元素保存到头结点的位置，
	h.tree[0] = save
	mpos := 0
	ipos := 0
	lpos := 0
	rpos := 0

	//校正，调整后满足堆的属性
	for {
		lpos = heapLeft(ipos)
		rpos = heapRight(ipos)

		//比较变化节点的左右孩子节点，判断变化的节点是否满足堆的性质(变化节点比左右孩子节点的值大)，
		//不满足，找到要调整的节点

		//取头节点和左孩子中较大的一个
		if lpos < h.size && h.compare(h.tree[lpos], h.tree[ipos]) > 0 {
			mpos = lpos
		} else {
			mpos = ipos
		}
		//比较上边的较大的和右结点的大小
		if rpos < h.size && h.compare(h.tree[rpos], h.tree[mpos]) > 0 {
			mpos = rpos
		}

		if mpos == ipos {
<<<<<<< HEAD
			//满足堆的性质，无需调整
			break
		} else {
			//调整位置，重新开始比较
=======
			//满足堆的性质,无需调整
			break
		} else {
			//调整位置,重新开始比较
>>>>>>> origin/master
			h.tree[mpos], h.tree[ipos] = h.tree[ipos], h.tree[mpos]
			ipos = mpos
		}

	}

	return data
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Top() interface{} {
	if h.size > 0 {
		return h.tree[0]
	}
	return nil
}

func (h *Heap) Data() []interface{} {
	return h.tree
}
