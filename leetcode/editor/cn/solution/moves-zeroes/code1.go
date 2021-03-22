package moves_zeroes

// 自己写的一种解法
// 思路：
// 当我发现一个零元素，我就跟我后面第一个非零元素进行交换。
func moveZeroes(n []int) {
	// 假设k个0，算法时间复杂度 O(n*k)
I:
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			for j := i + 1; j < len(n); j++ {
				if n[j] != 0 { // 当发现0后面第j个非0元素则交换
					n[i], n[j] = n[j], n[i]
					if j == len(n)-1 {
						break I
					}
					break
				}
			}
		}
	}
}
