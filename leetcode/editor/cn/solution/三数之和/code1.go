package 三数之和

import "sort"

// 只要做个排序
// 固定一个k元素，从k+1 和数组末尾夹逼
// 若三数之和大于零：动右指针
// 若三数之和小于零：动左指针
// 若三数之和等于零：加入结果，左右都动
func threeSum(nums []int) [][]int {
	three := make([][]int, 0)
	if len(nums) == 0 {
		return three
	}
	sort.Ints(nums)

	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break // 因为nums已经排好序了，当nums[k] >0 ，不可能出现等于0的情况
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		left, right := k+1, len(nums)-1
		for left < right {
			if nums[k]+nums[left]+nums[right] == 0 {
				three = append(three, []int{nums[k], nums[left], nums[right]})
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
			} else if nums[k]+nums[left]+nums[right] > 0 {
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else {
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
			}
		}
	}
	return three
}
