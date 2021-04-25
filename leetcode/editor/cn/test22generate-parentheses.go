// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
// 输入：n = 1
// 输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 回溯算法
// 👍 1717 👎 0

package main

import "fmt"

func main() {
	x := generateParenthesis(5)
	fmt.Println(x)
}

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	if n == 0 {
		return result
	}
	return dfs(result, "", 0, 0, n)
}

func dfs(result []string, curStr string, left, right int, n int) []string {
	if left == n && right == n {
		result = append(result, curStr)
		return result
	}

	// 剪支
	if left < right {
		return result
	}

	if left < n {
		result = dfs(result, curStr+"(", left+1, right, n)
	}
	if right < n {
		result = dfs(result, curStr+")", left, right+1, n)
	}
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
