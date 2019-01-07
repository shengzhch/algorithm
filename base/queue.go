package base

//队列 ：队头删除，对尾插入。先进先出。0表示对头元素下标

type Queue struct {
	values []ValueInterface
	rear   int
}

func NewQueue() *Queue {
	queue := &Queue{}
	queue.values = make([]ValueInterface, 0)
	queue.rear = -1

	return queue
}

func (q *Queue) Len() int {
	return q.rear + 1
}

func (q *Queue) IsEmpty() bool {
	return q.rear == -1
}


//出队列
func (q *Queue) Shift() (e ValueInterface) {
	e, q.values = q.values[0], q.values[1:]
	q.rear--
	return
}

//入队列
func (q *Queue) Push( v ValueInterface) {
	q.values = append(q.values, v)
	q.rear++
	return
}

func (q *Queue) Head() ValueInterface {
	return q.values[0]
}

func (q *Queue) Rear() ValueInterface {
	return q.values[q.rear]
}

