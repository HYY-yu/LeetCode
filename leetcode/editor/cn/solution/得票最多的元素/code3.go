package 得票最多的元素

// 分治法
func majorityElementF(nums []int) int {
	return mF(nums, 0, len(nums)-1)
}

func mF(nums []int, s, e int) int {
	// 终止条件
	if s == e {
		return nums[s]
	}

	mid := (e-s)/2 + s
	lnum := mF(nums, s, mid)
	rnum := mF(nums, mid+1, e)

	if lnum == rnum {
		return lnum
	}
	// 合并策略
	lcount := countNums(nums, lnum, s, e)
	rcount := countNums(nums, rnum, s, e)

	if lcount > rcount {
		return lnum
	}
	return rnum
}

func countNums(nums []int, num int, s, e int) int {
	count := 0
	for i := s; i <= e; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}
