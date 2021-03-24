package 最小的K个数

import "container/heap"

func getLeastNumbers(arr []int, k int) []int {
	intHeap := new(IntHeap2)

	heap.Init(intHeap)
	for i := 0; i < len(arr); i++ {
		heap.Push(intHeap, arr[i])
	}
	result := make([]int, k)
	for i := 0; i < len(result); i++ {
		result[i] = heap.Pop(intHeap).(int)
	}
	return result
}

type IntHeap2 []int

func (i IntHeap2) Len() int {
	return len(i)
}

func (h IntHeap2) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap2) Push(x interface{}) {
	if v, ok := x.(int); ok {
		*h = append(*h, v)
	}
}

func (h *IntHeap2) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}
