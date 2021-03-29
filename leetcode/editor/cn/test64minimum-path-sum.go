// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// Related Topics 数组 动态规划
// 👍 827 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
// 可否使用DFS做？

func minPathSum(grid [][]int) int {
	// DP[i][j] 指二维数组中某元素在此点的 最小路径
	// DP[i][j] = grid[i][j] + MIN(DP[i-1][j],DP[i][j-1])
	// DP[0][0] = grid[0][0]
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// 吸收大佬解法，不需要dp数组，直接在原grid上进行dp计算
	for i := 0; i < m; i++ {
		// 先特殊处理第一行
		if i == 0 {
			// 特殊处理第一列
			grid[i][0] = grid[i][0]
			for j := 1; j < n; j++ {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
		} else {
			// 特殊处理第一列
			grid[i][0] = grid[i][0] + grid[i-1][0]
			for j := 1; j < n; j++ {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

// 大佬解法
func minPathSum2(grid [][]int) int {
	// DP[i][j] 指二维数组中某元素在此点的 最小路径
	// DP[i][j] = grid[i][j] + MIN(DP[i-1][j],DP[i][j-1])
	// DP[0][0] = grid[0][0]
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)
