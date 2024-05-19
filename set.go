package go_algo

type Set[T comparable] map[T]bool

// NewSet 创建一个集合
func NewSet[T comparable]() *Set[T] {
	st := make(Set[T], 17)
	return &st
}

// Insert 往集合中新增一个元素
func (st *Set[T]) Insert(val T) {
	(*st)[val] = true
}

// Erase 删除集合中一个元素
func (st *Set[T]) Erase(val T) {
	delete(*st, val)
}

// Contains 判断集合中是否存在某个元素
func (st *Set[T]) Contains(val T) bool {
	_, exist := (*st)[val]
	return exist
}

// Size 求集合大小
func (st *Set[T]) Size() int {
	return len(*st)
}

// Empty 判断集合是否为空
func (st *Set[T]) Empty() bool {
	return len(*st) == 0
}

// Iterate 遍历集合，必须使用for-range遍历
func (st *Set[T]) Iterate() <-chan T {
	ch := make(chan T, 1)
	go func() {
		for k := range *st {
			ch <- k
		}
		close(ch)
	}()
	return ch
}
