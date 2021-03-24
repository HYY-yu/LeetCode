//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
// 示例 1：
//
// 输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//
//
// 示例 2：
//
// 输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//
//
// 限制：
//
//
// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000
//
// Related Topics 堆 分治算法
// 👍 217 👎 0

package main

import "container/heap"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)

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

//leetcode submit region end(Prohibit modification and deletion)
