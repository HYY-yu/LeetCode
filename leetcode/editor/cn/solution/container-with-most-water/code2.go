package container_with_most_water

import "math"

// i, j 往内部收缩
func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		minH := int(math.Min(float64(height[i]), float64(height[j])))
		area := (j - i) * minH
		max = int(math.Max(float64(max), float64(area)))
		if minH == height[i] {
			i++
		} else {
			j--
		}
	}
	return max
}
