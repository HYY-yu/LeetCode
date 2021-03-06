package 得票最多的元素

// 摩尔投票法
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cand := 0
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

	// 若无法保证必然存在 多于n/2 长度的元素，在选出cand后应该进行校验
	count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == cand {
			count++
		}
	}
	if count < len(nums)/2 {
		return 0
	}
	return cand
}

// CUT:  REVIEW TIME 2021.3.27 18:00
