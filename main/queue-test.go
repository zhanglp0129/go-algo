package main

import (
	"fmt"
	go_algo "go-algo"
)

func testBasicFunctionality() {
	queue := go_algo.NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if queue.Front() != 1 {
		fmt.Println("基本功能测试失败：队首元素错误")
	}

	if queue.Pop() != 1 {
		fmt.Println("基本功能测试失败：弹出元素错误")
	}

	if queue.Front() != 2 {
		fmt.Println("基本功能测试失败：队首元素错误")
	}

	queue.Push(4)
	queue.Push(5)

	if queue.Size() != 4 {
		fmt.Println("基本功能测试失败：队列大小错误")
	}
}

func testExpansion() {
	queue := go_algo.NewQueue[int]()
	for i := 0; i < 20; i++ {
		queue.Push(i)
	}

	if queue.Size() != 20 {
		fmt.Println("扩容测试失败：队列大小错误")
	}

	for i := 0; i < 20; i++ {
		if queue.Pop() != i {
			fmt.Println("扩容测试失败：元素错误")
		}
	}
}

func testCircularUsage() {
	queue := go_algo.NewQueue[int]()
	for i := 0; i < 10; i++ {
		queue.Push(i)
	}

	for i := 0; i < 5; i++ {
		if queue.Pop() != i {
			fmt.Println("循环使用测试失败：弹出元素错误")
		}
	}

	for i := 10; i < 15; i++ {
		queue.Push(i)
	}

	if queue.Size() != 10 {
		fmt.Println("循环使用测试失败：队列大小错误")
	}

	if queue.Front() != 5 {
		fmt.Println("循环使用测试失败：队首元素错误")
	}
}

func testEmptyQueue() {
	defer func() {
		if r := recover(); r == nil {
			fmt.Println("空队列测试失败：未触发panic")
		}
	}()

	queue := go_algo.NewQueue[int]()
	queue.Pop() // 应触发panic
}

func queueTest() {
	testBasicFunctionality()
	testExpansion()
	testCircularUsage()
	testEmptyQueue()
}
