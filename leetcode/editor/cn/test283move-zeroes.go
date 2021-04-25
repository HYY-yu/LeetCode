// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 980 👎 0

package main

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(n []int) {
	if len(n) <= 1 {
		return
	}
	j := 0
	for i := 0; i < len(n); i++ {
		if n[i] != 0 {
			n[i], n[j] = n[j], n[i]
			j++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
