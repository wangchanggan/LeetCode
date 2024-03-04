package double_pointer

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	r := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < k {
			r = i
			break
		}
	}
	if r == -1 || r == 0 {
		return 0
	}
	f := 0
	count := 0
	for f < r {
		if nums[f]+nums[r] == k {
			count++
			f++
			r--
		} else if nums[f]+nums[r] > k {
			r--
		} else {
			f++
		}
	}
	return count
}
