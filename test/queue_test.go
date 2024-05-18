package test

import (
	go_algo "go-algo"
	"testing"
)

func TestBasicFunctionality(t *testing.T) {
	queue := go_algo.NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if queue.Front() != 1 {
		t.Fatal("基本功能测试失败：队首元素错误")
	}

	if queue.Pop() != 1 {
		t.Fatal("基本功能测试失败：弹出元素错误")
	}

	if queue.Front() != 2 {
		t.Fatal("基本功能测试失败：队首元素错误")
	}

	queue.Push(4)
	queue.Push(5)

	if queue.Size() != 4 {
		t.Fatal("基本功能测试失败：队列大小错误")
	}
}

func TestExpansion(t *testing.T) {
	queue := go_algo.NewQueue[int]()
	for i := 0; i < 20; i++ {
		queue.Push(i)
	}

	if queue.Size() != 20 {
		t.Fatal("扩容测试失败：队列大小错误")
	}

	for i := 0; i < 20; i++ {
		if queue.Pop() != i {
			t.Fatal("扩容测试失败：元素错误")
		}
	}
}

func TestCircularUsage(t *testing.T) {
	queue := go_algo.NewQueue[int]()
	for i := 0; i < 10; i++ {
		queue.Push(i)
	}

	for i := 0; i < 5; i++ {
		if queue.Pop() != i {
			t.Fatal("循环使用测试失败：弹出元素错误")
		}
	}

	for i := 10; i < 15; i++ {
		queue.Push(i)
	}

	if queue.Size() != 10 {
		t.Fatal("循环使用测试失败：队列大小错误")
	}

	if queue.Front() != 5 {
		t.Fatal("循环使用测试失败：队首元素错误")
	}
}

func TestEmptyQueue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("空队列测试失败：未触发panic")
		}
	}()

	queue := go_algo.NewQueue[int]()
	queue.Pop() // 应触发panic
}
