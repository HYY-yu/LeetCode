package majority_element

// hash表记录
func majorityElementHash(nums []int) int {
	hash := make(map[int]int)
	n := len(nums)

	for i := 0; i < n; i++ {
		hash[nums[i]]++
		if hash[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}
