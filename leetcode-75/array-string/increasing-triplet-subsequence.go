package array_string

import "math"

func increasingTriplet(nums []int) bool {
	if nums == nil || len(nums) < 3 {
		return false
	}
	first := nums[0]
	second := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		if second < nums[i] {
			return true
		}
		if first < nums[i] {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}
