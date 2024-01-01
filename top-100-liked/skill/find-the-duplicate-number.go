package skill

import "sort"

func findDuplicate(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}
