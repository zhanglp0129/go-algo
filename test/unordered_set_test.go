package test

import (
	go_algo "go-algo"
	"testing"
)

func TestUnorderedSetInsertAndContains(t *testing.T) {
	set := go_algo.NewUnorderedSet[int]()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	if !set.Contains(1) {
		t.Fatal("集合应该包含元素 1")
	}
	if !set.Contains(2) {
		t.Fatal("集合应该包含元素 2")
	}
	if !set.Contains(3) {
		t.Fatal("集合应该包含元素 3")
	}
	if set.Contains(4) {
		t.Fatal("集合不应该包含元素 4")
	}
}

func TestUnorderedSetErase(t *testing.T) {
	set := go_algo.NewUnorderedSet[int]()
	set.Insert(1)
	set.Insert(2)
	set.Erase(1)

	if set.Contains(1) {
		t.Fatal("删除元素后，集合不应该再包含元素 1")
	}
	if !set.Contains(2) {
		t.Fatal("集合应该仍然包含元素 2")
	}
}

func TestUnorderedSetSize(t *testing.T) {
	set := go_algo.NewUnorderedSet[int]()
	if set.Size() != 0 {
		t.Fatal("新创建的集合应该是空的")
	}

	set.Insert(1)
	set.Insert(2)
	if set.Size() != 2 {
		t.Fatal("集合应该包含 2 个元素")
	}

	set.Insert(2) // 插入重复元素不应该改变大小
	if set.Size() != 2 {
		t.Fatal("当添加重复元素时，集合大小不应该增加")
	}

	set.Erase(1)
	if set.Size() != 1 {
		t.Fatal("删除一个元素后，集合大小应该减少到 1")
	}
}

func TestUnorderedSetEmpty(t *testing.T) {
	set := go_algo.NewUnorderedSet[int]()
	if !set.Empty() {
		t.Fatal("新创建的集合应该是空的")
	}

	set.Insert(1)
	if set.Empty() {
		t.Fatal("添加元素后，集合不应该是空的")
	}

	set.Erase(1)
	if !set.Empty() {
		t.Fatal("删除所有元素后，集合应该是空的")
	}
}

func TestUnorderedSetRemoveDuringIteration(t *testing.T) {
	set := go_algo.NewUnorderedSet[int]()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	// 遍历集合并尝试删除元素
	removed := false
	set.ToBegin()
	for set.Next() {
		if set.Get() == 2 && !removed {
			set.Erase(2)
			removed = true
		}
	}

	// 再次遍历检查元素 2 是否被正确删除
	elements := make(map[int]bool)
	set.ToBegin()
	for set.Next() {
		elements[set.Get()] = true
	}

	if elements[2] {
		t.Fatal("元素 2 应该已经从集合中删除")
	}
	if len(elements) != 2 {
		t.Fatal("删除元素 2 后，集合应该只包含 2 个元素")
	}
}
