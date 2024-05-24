package go_algo

type Queue[T any] struct {
	front, rear int // 最后一个元素不用，front->队首，rear->对尾的下一个位置
	data        []T // 数据
}

// NewQueue 创建队列
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{0, 0, make([]T, 15)}
}

// Push 添加一个元素到队尾
func (q *Queue[T]) Push(val T) {
	if q.Size()+1 == cap(q.data) { // 队列满了，需要扩容
		t := make([]T, 2*cap(q.data))
		size := cap(q.data) - 1
		for i := 0; i < size; i++ { // 将循环队列拉直
			t[i] = q.Pop()
		}
		q.front, q.rear, q.data = 0, size, t
	}
	q.data[q.rear] = val
	q.rear++
	if q.rear == cap(q.data) { // 队尾+1后，下标不能越界
		q.rear = 0
	}
}

// Pop 弹出队首元素，并返回
func (q *Queue[T]) Pop() T {
	t := q.Front() // 取出队首元素
	q.front++
	if q.front == cap(q.data) { // 队首+1后，下标不能越界
		q.front = 0
	}
	return t
}

// Front 返回队首元素
func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("不能从空队列中取元素")
	}
	return q.data[q.front]
}

// Size 求队列大小
func (q *Queue[T]) Size() int {
	size := q.rear - q.front
	if size < 0 { // 大小不能为负数
		return size + cap(q.data)
	}
	return size
}

// Empty 判断队列是否为空
func (q *Queue[T]) Empty() bool {
	return q.front == q.rear
}
