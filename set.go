package go_algo

import "cmp"

// Set 有序集合，基于红黑树实现
type Set[T any] struct {
	root *redBlackNode[T]  // 根节点
	comp func(a, b T) bool // 比较函数
}

// 定义颜色类型
type color uint32

const (
	red color = iota
	black
)

// RedBlackNode 红黑树结点
type redBlackNode[T any] struct {
	val         T
	left, right *redBlackNode[T]
	c           color
}

/* 针对红黑树的操作函数，均为不可导出 */

/* 针对有序集合的操作函数 */

// NewSetFunc 创建有序集合，需传入一个比较函数
func NewSetFunc[T any](comp func(a, b T) bool) *Set[T] {

}

// NewSet 创建有序集合，默认为升序
func NewSet[T cmp.Ordered]() *Set[T] {
	return NewSetFunc(func(a, b T) bool {
		return a < b
	})
}

func (st *Set[T]) Insert(val T) {

}

func (st *Set[T]) Erase(val T) {

}

func (st *Set[T]) Contains(val T) bool {

}

// TODO 实现UpperBound和LowerBound函数

func (st *Set[T]) Size() int {

}

func (st *Set[T]) Empty() bool {

}

func (st *Set[T]) ToBegin() {

}

func (st *Set[T]) Next() bool {

}

func (st *Set[T]) Get() T {

}
