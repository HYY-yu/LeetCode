package 移动零

// 1. 双指针法
// i 往前，发现不等于零，就和j交换
func moveZeros2(n []int) {
	// j
	// 1,3,4,0,0,1,2
	// i
	j := 0
	for i := 0; i < len(n); i++ {
		if n[i] != 0 {
			n[i], n[j] = n[j], n[i]
			j++
		}
	}
}
