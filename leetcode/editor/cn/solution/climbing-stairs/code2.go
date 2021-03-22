package climbing_stairs

// 不用dp数组的解法
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	p1, p2 := 0, 1
	for i := 0; i < n; i++ {
		p1, p2 = p2, p1+p2
	}
	return p2
}
