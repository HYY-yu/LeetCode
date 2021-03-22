package container_with_most_water

// i, j 双指针的另一种简洁写法
// 也是往内部收缩
func maxArea2(height []int) int {
	area, max := 0, 0
	i, j := 0, len(height)-1
	for i < j {
		if hi, hj, dist := height[i], height[j], j-i; hi > hj {
			area = dist * hj
			j--
		} else {
			area = dist * hi
			i++
		}

		if area > max {
			max = area
		}
	}
	return max
}
