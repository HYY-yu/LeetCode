package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 解题思路伪代码
	// n := 0
	// [0,0,0,1,1,1,2,2,3,3,4]
	//  i in (0,len(n))
	// a[i] != a[n] if i - 1 > n: swap(a[i],a[n+1]); n += 1
	// a[i] == a[n] next loop;
	// 优化
	// 1. 不需要swag，只要a[n+1] = a[i] 把i指向的值搞过去就行
	// 2. i 从 1 开始
	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[n] {
			if i-1 > n {
				nums[n+1] = nums[i]
			}
			n++
		}
	}
	return n + 1
}
