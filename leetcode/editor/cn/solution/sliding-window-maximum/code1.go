package sliding_window_maximum

import "container/heap"

// 最大堆实现滑动窗口
// 最大堆堆顶代表窗口中最大元素
// 若堆顶元素非窗口，应推出堆顶元素
// 也可用单调双端队列实现
type IntHeap [][]int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// 把元素推入int数组中
	if v, ok := x.([]int); ok {
		*h = append(*h, v)
	}
}

func (h *IntHeap) Pop() interface{} {
	o := *h
	n := len(o)
	v := o[n-1]
	*h = o[:n-1]
	return v
}

func (h IntHeap) Top() int {
	return h[0][0]
}

func (h IntHeap) TopIndex() int {
	return h[0][1]
}

func maxSlidingWindow(nums []int, k int) []int {
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	n := len(nums)
	result := make([]int, 0, n-k+1)

	for i := 0; i < k; i++ {
		heap.Push(intHeap, []int{nums[i], i})
	}

	// 插入首个元素
	result = append(result, intHeap.Top())

	for i := k; i < n; i++ {
		heap.Push(intHeap, []int{nums[i], i})
		// 如果堆顶元素超出窗口左边界，就应该弹出堆顶元素
		for intHeap.TopIndex() < i-k+1 {
			heap.Pop(intHeap)
		}
		result = append(result, intHeap.Top())
	}
	return result
}
