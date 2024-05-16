package go_algo

type Queue[T any] struct {
	front, rear int // 最后一个元素不用，front->队首，rear->对尾的下一个位置
	data        []T // 数据
}

// NewQueue 创建队列
func NewQueue[T any]() *Queue[T] {
	var q Queue[T]
	q.data = make([]T, 15) // 队列初始容量为15
	q.front, q.rear = 0, 0
	return &q
}

// Push 添加一个元素到队尾
func (q *Queue[T]) Push(val T) {
	if q.Size()+1 == cap(q.data) {
		t := make([]T, 2*cap(q.data))
		for i := 0; i < cap(q.data)-1; i++ {
			t[i] = q.Pop()
		}

		q.front, q.rear = 0, cap(q.data)-1
		q.data = t
	}
	q.data[q.rear] = val
	q.rear++
	if q.rear == cap(q.data) {
		q.rear = 0
	}
}

// Size 求队列大小
func (q *Queue[T]) Size() int {
	size := q.rear - q.front
	if size < 0 {
		return size + cap(q.data)
	}
	return size
}

// Empty 判断队列是否为空
func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

// Pop 弹出队首元素，并返回
func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("空队列不能弹出元素！")
	}
	t := q.Front()
	q.front++
	if q.front == cap(q.data) {
		q.front = 0
	}
	return t
}

// Front 返回队首元素
func (q *Queue[T]) Front() T {
	return q.data[q.front]
}
