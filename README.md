# go-algo
用Golang刷算法系列源码，提供了多种常用数据结构的实现模板，在做算法题时，直接复制到代码前即可

## 常用数据结构
### 队列
- 源码链接：[源码](queue.go)
- 文章链接：[文章](https://zhanglp.cn/archives/156)
- 创建队列：`func NewQueue[T any]() *Queue[T]`
- 入队：`func (q *Queue[T]) Push(val T)`
- 出队：`func (q *Queue[T]) Pop() T`
- 返回队首元素：`func (q *Queue[T]) Front() T`
- 求队列大小：`func (q *Queue[T]) Size() int`
- 判断队列是否为空：`func (q *Queue[T]) Empty() bool`

### 栈
- 创建栈：`func NewStack[T any]() *Stack[T]`
- 压栈：`func (stk *Stack[T]) Push(val T)`
- 弹栈：`func (stk *Stack[T]) Pop() T`
- 获取栈顶元素：`func (stk *Stack[T]) Top() T`
- 获取栈大小：`func (stk *Stack[T]) Size() int`
- 判断栈是否为空：`func (stk *Stack[T]) Empty() bool`

### Set集合
- 创建Set集合：`func NewSet[T comparable]() *Set[T]`
- 添加元素：`func (st *Set[T]) Insert(val T)`
- 删除元素：`func (st *Set[T]) Erase(val T)`
- 判断元素是否存在：`func (st *Set[T]) Contains(val T) bool`
- 求集合大小：`func (st *Set[T]) Size() int`
- 判断集合是否为空：`func (st *Set[T]) Empty() bool`
- 遍历集合：`func (st *Set[T]) Iterate() <-chan T`