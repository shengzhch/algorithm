/*
完全二叉树
叶子结点只在最后两层，最后一层结点从左到有排列
堆
子节点的值比父结点的值小（大），根结点是树中最大（小）的结点，称为最大（小）值堆。
用数组的方式存储数据，二叉树的层级遍历，左平衡树
*/

package basic

type Heap struct {
	size    int
	compare func(key1, key2 interface{}) int
	tree    []interface{}
}

func (h *Heap) Init(args ...interface{}) {
	h.size = 0
	h.tree = *new([]interface{})
	if len(args) == 1 {
		h.compare = args[0].(func(key1, key2 interface{}) int)
	}else{
		panic("compare func should be assign")
	}
	return
}

//结点下标的位置关系
var (
	//
	heap_parent = func(npos int) int {
		return (npos - 1) / 2
	}

	heap_left = func(npos int) int {
		return (npos * 2) + 1
	}

	heap_right = func(npos int) int {
		return (npos * 2) + 2
	}
)

//按最大堆实现

//插入 插入后保证堆的结构
func (h *Heap) Insert(data interface{}) {
	var ipos, ppos int

	h.tree = append(h.tree, data)

	ipos = h.size
	ppos = heap_parent(ipos)

	for (ipos > 0 && h.compare(h.tree[ppos], h.tree[ipos]) < 0) {
		h.tree[ppos], h.tree[ipos] = h.tree[ipos], h.tree[ppos]
		ipos = ppos
		ppos = heap_parent(ipos)
	}
	h.size++
}

//提取堆顶元素
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

	//最后一个元素保存到头结点的位置
	h.tree[0] = save
	mpos := 0
	ipos := 0
	lpos := 0
	rpos := 0

	//校正，调整后满足堆的属性
	for {
		lpos = heap_left(ipos)
		rpos = heap_right(ipos)


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

		if (mpos == ipos) {
			//满足堆的性质，无需调整
			break
		} else {
			//调整位置，重新开始比较
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

