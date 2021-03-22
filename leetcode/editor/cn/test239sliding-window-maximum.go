// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
// 。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
//
// 示例 3：
//
//
// 输入：nums = [1,-1], k = 1
// 输出：[1,-1]
//
//
// 示例 4：
//
//
// 输入：nums = [9,11], k = 2
// 输出：[11]
//
//
// 示例 5：
//
//
// 输入：nums = [4,-2], k = 2
// 输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window
// 👍 888 👎 0

package main

import "container/heap"

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)
