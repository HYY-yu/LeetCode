package 加1

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}

// 加k 思路
