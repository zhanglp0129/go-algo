# go-algo
1. 用Golang刷算法系列源码，提供了多种常用数据结构的实现模板，在做算法题时，直接复制到代码前即可
2. 代码风格与C++的STL基本相同，只是进行少数修改

## 常用数据结构
### 队列
- [源码](queue.go)
- 文章链接：[文章](https://zhanglp.cn/archives/156)
- 创建队列：`func NewQueue[T any]() *Queue[T]`
- 入队：`func (q *Queue[T]) Push(val T)`
- 出队：`func (q *Queue[T]) Pop() T`
- 返回队首元素：`func (q *Queue[T]) Front() T`
- 求队列大小：`func (q *Queue[T]) Size() int`
- 判断队列是否为空：`func (q *Queue[T]) Empty() bool`

### 栈
- [源码](stack.go)
- 文章链接：[文章](https://zhanglp.cn/archives/170)
- 创建栈：`func NewStack[T any]() *Stack[T]`
- 压栈：`func (stk *Stack[T]) Push(val T)`
- 弹栈：`func (stk *Stack[T]) Pop() T`
- 获取栈顶元素：`func (stk *Stack[T]) Top() T`
- 获取栈大小：`func (stk *Stack[T]) Size() int`
- 判断栈是否为空：`func (stk *Stack[T]) Empty() bool`

### 优先队列
- [源码](priority_queue.go)
- 创建优先队列，默认为大根堆：`func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T]`
- 创建优先队列，\<表示大根堆，\>表示小根堆：`func NewPriorityQueueFunc[T any](comp func(a, b T) bool) *PriorityQueue[T]`
- 入队：`func (pq *PriorityQueue[T]) Push(val T)`
- 出队：`func (pq *PriorityQueue[T]) Pop() T`
- 获取堆顶元素：`func (pq *PriorityQueue[T]) Top() T`
- 求优先队列大小：`func (pq *PriorityQueue[T]) Size() int`
- 判断优先队列是否为空：`func (pq *PriorityQueue[T]) Empty() bool`

### 无序集合
- [源码](unordered_set.go)
- 创建无序集合，仅支持整数类型：`func NewUnorderedSet[T number]() *UnorderedSet[T]`
- 创建无序集合，需传入哈希函数和相等判断函数：`func NewUnorderedSetFunc[T any](hash func(a T) uint64, equal func(a, b T) bool) *UnorderedSet[T]`
- 插入一个元素：`func (us *UnorderedSet[T]) Insert(val T)`
- 删除一个元素：`func (us *UnorderedSet[T]) Erase(val T)`
- 判断集合中是否存在该元素：`func (us *UnorderedSet[T]) Contains(val T) bool`
- 求集合大小：`func (us *UnorderedSet[T]) Size() int`
- 判断集合是否为空：`func (us *UnorderedSet[T]) Empty() bool`
- 遍历集合：
    - 将内置迭代器指向初始位置：`func (us *UnorderedSet[T]) ToBegin()`
    - 将内置迭代器指向下一个位置，并返回是否成功：`func (us *UnorderedSet[T]) Next() bool`
    - 获取当前迭代器指向的元素：`func (us *UnorderedSet[T]) Get() T`
    - 示例：
    ```go
    us := NewUnorderedSet[int]()
    us.ToBegin()    // 遍历前先复位
    for us.Next() {
        v := us.Get()   // 获取当前遍历的元素
        fmt.Println(v)
    }
    ```

    