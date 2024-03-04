package sliding_window

import "math"

func findMaxAverage(nums []int, k int) float64 {
	if nums == nil {
		return 0
	} else if len(nums) == k {
		var sum int
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
		}
		return float64(sum) / float64(k)
	}

	max := -math.MaxFloat64
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	tmp := float64(sum) / float64(k)
	if tmp > max {
		max = tmp
	}
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		tmp := float64(sum) / float64(k)
		if tmp > max {
			max = tmp
		}
	}
	return max
}
