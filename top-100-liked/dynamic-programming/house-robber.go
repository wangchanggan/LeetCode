package dynamic_programming

func rob(nums []int) int {
	if nums == nil {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return maxThree(nums[0], nums[1], 0)
	} else if len(nums) == 3 {
		return maxThree(nums[0], nums[1], nums[0]+nums[2])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = maxThree(nums[0], nums[1], nums[0]+nums[2])
	for i := 3; i < len(nums); i++ {
		dp[i] = maxThree(dp[i-1], nums[i]+dp[i-2], nums[i]+dp[i-3])
	}
	return dp[len(dp)-1]
}

func maxThree(a, b, c int) int {
	if a > b {
		if c > a {
			return c
		} else {
			return a
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
