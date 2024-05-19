package problems

// 直接复制队列源码
type Queue[T any] struct {
	front, rear int
	data        []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{0, 0, make([]T, 15)}
}
func (q *Queue[T]) Push(val T) {
	if q.Size()+1 == cap(q.data) {
		t := make([]T, 2*cap(q.data))
		size := cap(q.data) - 1
		for i := 0; i < size; i++ {
			t[i] = q.Pop()
		}
		q.front, q.rear, q.data = 0, size, t
	}
	q.data[q.rear] = val
	q.rear++
	if q.rear == cap(q.data) {
		q.rear = 0
	}
}
func (q *Queue[T]) Pop() T {
	t := q.Front()
	q.front++
	if q.front == cap(q.data) {
		q.front = 0
	}
	return t
}
func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("不能从空队列中取元素")
	}
	return q.data[q.front]
}
func (q *Queue[T]) Size() int {
	size := q.rear - q.front
	if size < 0 {
		return size + cap(q.data)
	}
	return size
}
func (q *Queue[T]) Empty() bool {
	return q.front == q.rear
}

// 真正的解题代码
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res, count := -1, 0 // count 新鲜橘子的数量
	// 定义结构体，表示一个单元格的坐标
	type pair struct {
		x, y int
	}
	q := NewQueue[pair]()
	// 4个方向
	ways := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 往队列中插入初始数据
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				q.Push(pair{i, j})
			}
		}
	}

	// 没有橘子
	if q.Size() == 0 && count == 0 {
		return 0
	}

	// 广度优先搜索
	for !q.Empty() {
		size := q.Size()
		for size > 0 {
			p := q.Pop()
			size--
			x, y := p.x, p.y

			// 遍历4个方向
			for _, way := range ways {
				dx, dy := way[0], way[1]
				nx, ny := dx+x, dy+y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					q.Push(pair{nx, ny})
					grid[nx][ny] = 2
					count--
				}
			}
		}
		res++
	}

	// 没有全部腐烂
	if count > 0 {
		return -1
	}
	return res
}
