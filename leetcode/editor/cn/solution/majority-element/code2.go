package majority_element

// 摩尔投票法
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cand := nums[0]
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			cand = nums[i]
		}
		if cand == nums[i] {
			count++
		} else {
			count--
		}
	}

	// 若无法保证必然存在 多余n/2 长度的元素，在选出cand后应该进行校验
	return cand
}
