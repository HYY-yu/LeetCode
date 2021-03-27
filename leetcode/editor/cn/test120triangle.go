// 给定一个三角形 三角形的最小路径和 ，找出自顶向下的最小路径和。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果
// 正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
//
//
// 示例 1：
//
//
// 输入：三角形的最小路径和 = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//    2
//   3 4
//  6 5 7
// 4 1 8 3
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
// f(i,j)=t[i][j] + min(f(i+1,j),f(i+1,j+1))
// f(0,0) = 答案
// 示例 2：
//
//
// 输入：三角形的最小路径和 = [[-10]]
// 输出：-10
//
//
//
//
// 提示：
//
//
// 1 <= 三角形的最小路径和.length <= 200
// 三角形的最小路径和[0].length == 1
// 三角形的最小路径和[i].length == 三角形的最小路径和[i - 1].length + 1
// -104 <= 三角形的最小路径和[i][j] <= 104
//

//
//
// 进阶：
//
//
// 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
//
// Related Topics 数组 动态规划
// 👍 692 👎 0

// f(i-1,-1) = 0 ,x in ALL
// f(i-1, j) = 0 ,j==t[i].len-1
// f(0,0) = t[0][0]
// f(1,0) = t[1][0] + min(f(0,-1) , f(0,0))
// f(1,1) = t[1][1] + min((0,0) , f(0,1))
// f(i,j) = t[i][j] + min(f(i-1,j-1) , f(i-1,j))
package main

import (
	"fmt"
	"math"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}

	r := minimumTotal(triangle)
	fmt.Println(r)
	fmt.Println(r == 11)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	deepIndex := len(triangle)
	res := make([]int, deepIndex)
	copy(res, triangle[deepIndex-1])

	for i := deepIndex - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[j] = triangle[i][j] + int(math.Min(float64(res[j+1]), float64(res[j])))
		}
	}
	return res[0]
}

// leetcode submit region end(Prohibit modification and deletion)
