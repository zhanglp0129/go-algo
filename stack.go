package go_algo

type Stack[T any] []T

// NewStack 初始化栈
func NewStack[T any]() *Stack[T] {
	stk := make(Stack[T], 0, 15)
	return &stk
}

// Push 压栈
func (stk *Stack[T]) Push(val T) {
	*stk = append(*stk, val)
}

// Pop 弹栈
func (stk *Stack[T]) Pop() T {
	if stk.Empty() {
		panic("空栈不能弹出元素！")
	}
	t := (*stk)[len(*stk)-1]
	*stk = (*stk)[:len(*stk)-1]
	return t
}

// Top 获取栈顶元素
func (stk *Stack[T]) Top() T {
	return (*stk)[len(*stk)-1]
}

// Size 获取栈大小
func (stk *Stack[T]) Size() int {
	return len(*stk)
}

// Empty 判断栈是否为空
func (stk *Stack[T]) Empty() bool {
	return stk.Size() == 0
}
