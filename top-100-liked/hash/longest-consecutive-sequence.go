package hash

import "sort"

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			count++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}
