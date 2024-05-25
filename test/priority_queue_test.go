package test

import (
	go_algo "go-algo"
	"math/rand"
	"slices"
	"testing"
)

// 测试用优先队列实现堆排序
func TestHeapSort(t *testing.T) {
	// 创建小根堆
	pq := go_algo.NewPriorityQueueFunc[int](func(a, b int) bool {
		return a > b
	})

	// 生成100个随机数，并插入堆
	for i := 0; i < 100; i++ {
		pq.Push(rand.Int())
	}

	// 堆排序
	s := make([]int, 0, 100)
	for !pq.Empty() {
		s = append(s, pq.Pop())
		if len(s)+pq.Size() != 100 {
			t.Fatal("堆排序测试失败，优先队列中元素个数不正确")
		}
	}

	// 验证有序性
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			t.Fatal("堆排序测试失败，排序后未保证有序性")
		}
	}
}

// 测试优先队列随机插入删除
func TestPriorityQueueRandomPushPop(t *testing.T) {
	// 创建大根堆
	pq := go_algo.NewPriorityQueue[int]()

	in := make([]int, 0) // 优先队列中的所有元素

	// 随机插入删除，最多执行100次
	i := 0
	for i < 100 || !pq.Empty() {
		if rand.Int()%2 == 1 {
			// 插入
			r := rand.Int()
			pq.Push(r)
			in = append(in, r)
			i++
		} else if !pq.Empty() {
			// 删除并判断该元素是否最大
			top := pq.Pop() // 堆顶元素
			slices.Sort(in)
			if top != in[len(in)-1] {
				t.Fatal("测试优先队列随机插入删除失败，堆顶元素不是最大")
			}
			in = in[:len(in)-1]
		}

		// 测试长度
		if len(in) != pq.Size() {
			t.Fatal("测试优先队列随机插入删除失败，优先队列中元素个数不正确")
		}
	}
}
