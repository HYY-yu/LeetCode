package 得票最多的元素

import "math"

func majorityElementMore(nums []int) []int {
	M := 2 // 候选人数量
	if len(nums) == 0 {
		return nil
	}
	cand := make([]int, M)
	count := make([]int, M)

	// 投票阶段
I:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(cand); j++ {

			if nums[i] == cand[j] {
				// 投票
				count[j]++
				continue I
			}
		}
		for j := 0; j < len(cand); j++ {
			if count[j] == 0 {
				// 换候选人重新计票
				cand[j] = nums[i]
				count[j]++
				continue I
			}
		}
		// 消票
		for j := 0; j < len(cand); j++ {
			count[j]--
		}
	}

	// 计票阶段
	total := make([]int, M)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(total); j++ {
			if nums[i] == cand[j] {
				total[j]++
			}
		}
	}
	up := math.MaxInt64 // 对 0,0,0 做特殊处理
	result := make([]int, 0, M)
	for i := 0; i < len(total); i++ {
		if total[i] > len(nums)/(M+1) && up != 0 {
			result = append(result, cand[i])
			up = cand[i]
		}
	}
	return result
}
