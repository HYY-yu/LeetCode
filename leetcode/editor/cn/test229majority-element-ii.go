// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
//
//
//
// 示例 1：
//
//
// 输入：[3,2,3]
// 输出：[3]
//
// 示例 2：
//
//
// 输入：nums = [1]
// 输出：[1]
//
//
// 示例 3：
//
//
// 输入：[1,1,1,3,3,2,2,2]
// 输出：[1,2]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109
//
// Related Topics 数组
// 👍 351 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{2, 1, 1, 3, 1, 4, 5, 6}
	x := majorityElementMore(a)
	fmt.Println(x)
}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElementMore(nums []int) []int {
	M := 2 // 候选人数量
	if len(nums) == 0 {
		return nil
	}
	cand := make([]int, M)
	count := make([]int, M)

	// 投票阶段
I:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(cand); j++ {

			if nums[i] == cand[j] {
				// 投票
				count[j]++
				continue I
			}
		}
		for j := 0; j < len(cand); j++ {
			if count[j] == 0 {
				// 换候选人重新计票
				cand[j] = nums[i]
				count[j]++
				continue I
			}
		}
		// 消票
		for j := 0; j < len(cand); j++ {
			count[j]--
		}
	}

	// 计票阶段
	total := make([]int, M)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(total); j++ {
			if nums[i] == cand[j] {
				total[j]++
			}
		}
	}
	up := math.MaxInt64 // 对 0,0,0 做特殊处理
	result := make([]int, 0, M)
	for i := 0; i < len(total); i++ {
		if total[i] > len(nums)/(M+1) && up != 0 {
			result = append(result, cand[i])
			up = cand[i]
		}
	}
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
