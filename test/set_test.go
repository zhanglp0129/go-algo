package test

import (
	go_algo "go-algo"
	"testing"
)

// TestSetInsert 测试 Insert 和 Contains 方法
func TestSetInsert(t *testing.T) {
	set := go_algo.NewSet[int]()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	if !set.Contains(1) {
		t.Fatal("预期集合包含 1")
	}
	if !set.Contains(2) {
		t.Fatal("预期集合包含 2")
	}
	if !set.Contains(3) {
		t.Fatal("预期集合包含 3")
	}
}

// TestSetSize 测试 Size 方法
func TestSetSize(t *testing.T) {
	set := go_algo.NewSet[int]()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	if set.Size() != 3 {
		t.Fatalf("预期集合大小为 3，得到 %d", set.Size())
	}
}

// TestSetErase 测试 Erase 和 Contains 方法
func TestSetErase(t *testing.T) {
	set := go_algo.NewSet[int]()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	set.Erase(2)
	if set.Contains(2) {
		t.Fatal("预期集合不包含 2")
	}
	if set.Size() != 2 {
		t.Fatalf("预期集合大小为 2，得到 %d", set.Size())
	}
}

// TestSetEmpty 测试 Empty 方法
func TestSetEmpty(t *testing.T) {
	set := go_algo.NewSet[int]()
	if !set.Empty() {
		t.Fatal("预期集合为空")
	}

	set.Insert(1)
	if set.Empty() {
		t.Fatal("预期集合不为空")
	}

	set.Erase(1)
	if !set.Empty() {
		t.Fatal("预期集合在删除最后一个元素后为空")
	}
}

// TestSetIterate 测试 Iterate 方法
func TestSetIterate(t *testing.T) {
	set := go_algo.NewSet[int]()
	set.Insert(4)
	set.Insert(5)
	set.Insert(6)

	expected := map[int]bool{4: true, 5: true, 6: true}
	for val := range set.Iterate() {
		if !expected[val] {
			t.Fatalf("遍历时遇到意外值: %d", val)
		}
		delete(expected, val)
	}
	if len(expected) != 0 {
		t.Fatal("并非所有预期值都被遍历")
	}
}
