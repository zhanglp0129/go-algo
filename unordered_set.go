package go_algo

import (
	"math/rand"
)

// UnorderedSet 无序集合
type UnorderedSet[T any] struct {
	bucket   []*listNode[T]    // 存储桶
	size     int               // 集合大小
	hash     func(a T) uint64  // 哈希函数
	equal    func(a, b T) bool // 相等判断函数
	curIndex int               // 迭代器指向的元素下标
	curList  *listNode[T]      // 迭代器指向的链表结点
}

// 链表结点
type listNode[T any] struct {
	data T
	next *listNode[T]
}

type number interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 |
		~int32 | ~uint32 | ~int64 | ~uint64 | ~uintptr
}

// NewUnorderedSet 创建无序集合，为整型提供默认的哈希函数和相等判断函数
func NewUnorderedSet[T number]() *UnorderedSet[T] {
	hash := func(a T) uint64 {
		return uint64(a)
	}
	equal := func(a, b T) bool {
		return a == b
	}
	return NewUnorderedSetFunc(hash, equal)
}

// NewUnorderedSetFunc 创建无序集合，需传入hash函数和相等判断函数
func NewUnorderedSetFunc[T any](hash func(a T) uint64, equal func(a, b T) bool) *UnorderedSet[T] {
	return &UnorderedSet[T]{
		bucket:   make([]*listNode[T], 17),
		size:     0,
		hash:     hash,
		equal:    equal,
		curIndex: -1,
		curList:  nil,
	}
}

// Insert 插入一个元素
func (us *UnorderedSet[T]) Insert(val T) {
	if us.Contains(val) { // 先检查是否存在该元素
		return
	}

	if us.size == len(us.bucket) {
		// 扩容，长度为原来的2倍+1或-1
		var newBucket []*listNode[T]
		if rand.Int()%2 == 1 {
			newBucket = make([]*listNode[T], 2*us.size+1)
		} else {
			newBucket = make([]*listNode[T], 2*us.size-1)
		}

		// 拷贝到newBucket
		us.ToBegin()    // 先复位迭代器
		for us.Next() { // 遍历所有元素
			us.insertIntoBucket(newBucket, us.Get())
		}
		us.ToBegin()
		us.bucket = newBucket
	}

	// 插入元素
	us.insertIntoBucket(us.bucket, val)
	us.size++
}

// 中间函数：往指定存储桶中插入一个元素
func (us *UnorderedSet[T]) insertIntoBucket(bucket []*listNode[T], val T) {
	idx := us.hash(val) % uint64(len(bucket))
	if bucket[idx] == nil { // 头节点为空
		bucket[idx] = &listNode[T]{val, nil}
	} else { // 头节点不为空
		head := bucket[idx]
		for head.next != nil {
			head = head.next
		}
		head.next = &listNode[T]{val, nil}
	}
}

// Erase 删除一个元素
func (us *UnorderedSet[T]) Erase(val T) {
	if !us.Contains(val) {
		return
	}

	idx := us.hash(val) % uint64(len(us.bucket))
	if us.equal(us.bucket[idx].data, val) {
		us.bucket[idx] = us.bucket[idx].next
	} else {
		head := us.bucket[idx]
		for !us.equal(head.next.data, val) {
			head = head.next
		}
		head.next = head.next.next
	}
	us.size--
}

// Contains 判断集合中是否存在该元素
func (us *UnorderedSet[T]) Contains(val T) bool {
	idx := us.hash(val) % uint64(len(us.bucket))
	head := us.bucket[idx]
	for head != nil {
		if us.equal(head.data, val) {
			return true
		}
		head = head.next
	}
	return false
}

// Size 求集合大小
func (us *UnorderedSet[T]) Size() int {
	return us.size
}

// Empty 判断集合是否为空
func (us *UnorderedSet[T]) Empty() bool {
	return us.size == 0
}

// ToBegin 将内置迭代器指向初始位置，每次遍历前记得调用该方法
func (us *UnorderedSet[T]) ToBegin() {
	us.curIndex, us.curList = -1, nil
}

// Next 将内置迭代器指向下一个位置，并返回是否成功
func (us *UnorderedSet[T]) Next() bool {
	if us.curList != nil {
		us.curList = us.curList.next
	}

	for us.curList == nil {
		us.curIndex++
		if us.curIndex == len(us.bucket) {
			return false
		}
		us.curList = us.bucket[us.curIndex]
	}
	return true
}

// Get 获取当前迭代器指向的元素
func (us *UnorderedSet[T]) Get() T {
	if us.curList == nil {
		panic("内置迭代器没有指向合法的值")
	}
	return us.curList.data
}
