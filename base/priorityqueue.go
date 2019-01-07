package base

//优先队列： 每次从队列中取出的元素都是最高优先权的元素
//存储表示和实现方法有多种，可用数组或者堆来实现，该文件采用数组的方式 堆结构单独介绍

//查找 插入 删除

//最小优先队列 查找优先权最小的元素 删除该元素
//最大优先队列 查找优先权最大的元素 删除该元素

//优先队列中的元素
type PValueInterface interface {
	ValueInterface
	Priority(PValueInterface) bool
}

const (
	minpq = iota
)

//下标为0的优先级最大
type PQueue struct {
	values []PValueInterface
	top    int
}

func NewPQueue() *PQueue {
	return &PQueue{
		values: make([]PValueInterface, 0),
		top:    -1,
	}
}

func (pq *PQueue) Find(kind int) PValueInterface {
	if pq.top == -1 {
		return nil
	}
	if kind == minpq {
		return pq.values[pq.top]
	}
	return pq.values[0]
}

func (pq *PQueue) Insert(pv PValueInterface) {
	for i, v := range pq.values {
		if i == pq.top {
			pq.values = append(pq.values, pv)
			pq.top++
			return
		}
		vNext := pq.values[i+1]
		if v.Priority(pv) && (!vNext.Priority(pv)) {
			tp := append(pq.values[:i+1], pv)
			pq.values = append(tp, pq.values[i+1:]...)
			pq.top++
			return
		}
	}
	return
}

func (pq *PQueue) Remove(kind int) {
	if pq.top == -1 {
		return
	}
	if kind == minpq {
		pq.values = pq.values[:pq.top]
	} else {
		pq.values = pq.values[1:]
	}
	pq.top --
	return
}
