package main

import (
	"errors"
	"fmt"
	go_algo "go-algo"
)

func testPushAndPop() error {
	stk := go_algo.NewStack[int]()
	stk.Push(1)
	stk.Push(2)
	if stk.Pop() != 2 {
		return errors.New("测试Push和Pop失败：期望2，得到" + fmt.Sprint(stk.Top()))
	}
	if stk.Pop() != 1 {
		return errors.New("测试Push和Pop失败：期望1，得到" + fmt.Sprint(stk.Top()))
	}
	if !stk.Empty() {
		return errors.New("测试Push和Pop失败：期望栈为空")
	}
	return nil
}

func testTop() error {
	stk := go_algo.NewStack[int]()
	stk.Push(1)
	if stk.Top() != 1 {
		return errors.New("测试Top失败：期望1，得到" + fmt.Sprint(stk.Top()))
	}
	stk.Push(2)
	if stk.Top() != 2 {
		return errors.New("测试Top失败：期望2，得到" + fmt.Sprint(stk.Top()))
	}
	return nil
}

func testSize() error {
	stk := go_algo.NewStack[int]()
	if stk.Size() != 0 {
		return errors.New("测试Size失败：期望0，得到" + fmt.Sprint(stk.Size()))
	}
	stk.Push(1)
	stk.Push(2)
	if stk.Size() != 2 {
		return errors.New("测试Size失败：期望2，得到" + fmt.Sprint(stk.Size()))
	}
	stk.Pop()
	if stk.Size() != 1 {
		return errors.New("测试Size失败：期望1，得到" + fmt.Sprint(stk.Size()))
	}
	return nil
}

func testEmpty() error {
	stk := go_algo.NewStack[int]()
	if !stk.Empty() {
		return errors.New("测试Empty失败：期望栈为空")
	}
	stk.Push(1)
	if stk.Empty() {
		return errors.New("测试Empty失败：期望栈不为空")
	}
	stk.Pop()
	if !stk.Empty() {
		return errors.New("测试Empty失败：期望栈为空")
	}
	return nil
}

func stackTest() {
	tests := []struct {
		name string
		test func() error
	}{
		{"测试Push和Pop", testPushAndPop},
		{"测试Top", testTop},
		{"测试Size", testSize},
		{"测试Empty", testEmpty},
	}

	for _, t := range tests {
		if err := t.test(); err != nil {
			fmt.Println(t.name, "失败：", err)
		}
	}
}
