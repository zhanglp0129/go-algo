package go_algo

import (
	"cmp"
)

type PriorityQueue[T any] struct {
	c    []T               // 存储数据
	comp func(a, b T) bool // 比较函数。满足：comp(子结点, 父结点)==true；右边的元素优先级更高
}

// NewPriorityQueue 创建一个优先队列，默认为大根堆
func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return NewPriorityQueueFunc(func(a, b T) bool {
		return a < b // <表示大根堆
	})
}

// NewPriorityQueueFunc 创建一个优先队列，<表示大根堆，>表示小根堆
func NewPriorityQueueFunc[T any](comp func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		c:    make([]T, 0, 15),
		comp: comp,
	}
}

// Push 往优先队列插入一个元素
func (pq *PriorityQueue[T]) Push(val T) {
	// 满足：子结点=i，父结点=(i-1)/2
	pq.c = append(pq.c, val)
	for i := len(pq.c) - 1; i > 0 && pq.comp(pq.c[(i-1)/2], pq.c[i]); i = (i - 1) / 2 {
		// 如果父子结点关系不满足comp(子结点, 父结点)，则交换
		pq.c[i], pq.c[(i-1)/2] = pq.c[(i-1)/2], pq.c[i]
	}
}

// Pop 从优先队列移除一个元素并返回
func (pq *PriorityQueue[T]) Pop() T {
	v := pq.Top() // 返回值
	// 交换第一个元素和最后一个元素，然后删除最后一个元素
	pq.c[0], pq.c[len(pq.c)-1] = pq.c[len(pq.c)-1], pq.c[0]
	pq.c = pq.c[:len(pq.c)-1]
	// 下沉第一个元素，左子结点=2*i+1，右子结点=2*i+2，父结点=i
	i := 0
	for 2*i+1 < len(pq.c) {
		j := 2*i + 1 // j指向优先级更高的子结点
		if 2*i+2 < len(pq.c) && pq.comp(pq.c[2*i+1], pq.c[2*i+2]) {
			j = 2*i + 2 // 2*i+2优先级更高
		}

		if pq.comp(pq.c[i], pq.c[j]) { // 父子结点关系不满足comp(子结点, 父结点)
			pq.c[i], pq.c[j] = pq.c[j], pq.c[i]
		} else {
			break
		}
		i = j
	}
	return v
}

// Top 返回第一个元素
func (pq *PriorityQueue[T]) Top() T {
	if pq.Empty() {
		panic("空优先队列不能取元素")
	}
	return pq.c[0]
}

// Size 求优先队列大小
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.c)
}

// Empty 判断优先队列是否为空
func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.c) == 0
}
