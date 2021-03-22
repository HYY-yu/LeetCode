// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
// 复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
// 输入：nums = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：nums = [0]
// 输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// Related Topics 数组 双指针
// 👍 3056 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-2, 0, 0, 2, 2}
	r := threeSum(a)
	fmt.Println(r)
}

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	three := make([][]int, 0)
	if len(nums) == 0 {
		return three
	}
	sort.Ints(nums)

	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break // 因为nums已经排好序了，当nums[k] >0 ，不可能出现等于0的情况
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		left, right := k+1, len(nums)-1
		for left < right {
			if nums[k]+nums[left]+nums[right] == 0 {
				three = append(three, []int{nums[k], nums[left], nums[right]})
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
			} else if nums[k]+nums[left]+nums[right] > 0 {
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else {
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
			}
		}
	}
	return three
}

// leetcode submit region end(Prohibit modification and deletion)
