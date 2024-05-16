# go-algo
用Golang刷算法系列源码，提供了多种常用数据结构的实现模板，在做算法题时，直接复制到代码前即可

## 常用数据结构
### 队列
- 代码文件：queue.go
- 创建队列：`func NewQueue[T any]() *Queue[T]`
- 入队：`func (q *Queue[T]) Push(val T)`
- 出队：`func (q *Queue[T]) Pop() T`
- 返回队首元素：`func (q *Queue[T]) Front() T`
- 求队列大小：`func (q *Queue[T]) Size() int`
- 判断队列是否为空：`func (q *Queue[T]) Empty() bool`