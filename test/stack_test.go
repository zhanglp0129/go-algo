package test

import (
	"fmt"
	go_algo "go-algo"
	"testing"
)

func TestPushAndPop(t *testing.T) {
	stk := go_algo.NewStack[int]()
	stk.Push(1)
	stk.Push(2)
	if stk.Pop() != 2 {
		t.Fatal("测试Push和Pop失败：期望2，得到" + fmt.Sprint(stk.Top()))
	}
	if stk.Pop() != 1 {
		t.Fatal("测试Push和Pop失败：期望1，得到" + fmt.Sprint(stk.Top()))
	}
	if !stk.Empty() {
		t.Fatal("测试Push和Pop失败：期望栈为空")
	}
}

func TestTop(t *testing.T) {
	stk := go_algo.NewStack[int]()
	stk.Push(1)
	if stk.Top() != 1 {
		t.Fatal("测试Top失败：期望1，得到" + fmt.Sprint(stk.Top()))
	}
	stk.Push(2)
	if stk.Top() != 2 {
		t.Fatal("测试Top失败：期望2，得到" + fmt.Sprint(stk.Top()))
	}
}

func TestSize(t *testing.T) {
	stk := go_algo.NewStack[int]()
	if stk.Size() != 0 {
		t.Fatal("测试Size失败：期望0，得到" + fmt.Sprint(stk.Size()))
	}
	stk.Push(1)
	stk.Push(2)
	if stk.Size() != 2 {
		t.Fatal("测试Size失败：期望2，得到" + fmt.Sprint(stk.Size()))
	}
	stk.Pop()
	if stk.Size() != 1 {
		t.Fatal("测试Size失败：期望1，得到" + fmt.Sprint(stk.Size()))
	}
}

func TestEmpty(t *testing.T) {
	stk := go_algo.NewStack[int]()
	if !stk.Empty() {
		t.Fatal("测试Empty失败：期望栈为空")
	}
	stk.Push(1)
	if stk.Empty() {
		t.Fatal("测试Empty失败：期望栈不为空")
	}
	stk.Pop()
	if !stk.Empty() {
		t.Fatal("测试Empty失败：期望栈为空")
	}
}
