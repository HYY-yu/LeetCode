package plus_k

import "math"

// 某数加k 思路
func plusK(digits []int, k int) []int {
	// 转化成数字，再转化回来
	return Int2IntArr(IntArr2Int(digits) + k)
}

func IntArr2Int(d []int) int {
	var result int
	n := len(d) - 1
	for i := n; i >= 0; i-- {
		pow := math.Pow10(n - i)
		result += d[i] * int(pow)
	}
	return result
}

func Int2IntArr(t int) []int {
	result := make([]int, 0)
	n := 1
	for {
		pow := int(math.Pow10(n))
		y := t % pow
		x := y / (pow / 10)
		if x == 0 {
			result = append(result, y)
		} else {
			result = append(result, x)
		}
		if y == t {
			break
		}
		n++
	}

	lo, hi := 0, len(result)-1
	for lo < hi {
		result[lo], result[hi] = result[hi], result[lo]
		lo++
		hi--
	}
	return result
}
