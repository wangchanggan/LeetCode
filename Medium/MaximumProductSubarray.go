/*Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.*/

package Medium

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		curMax, curMin = curMax*num, curMin*num
		curMax, curMin = max(curMax, curMin, num), min(curMax, curMin, num)
		res = int(math.Max(float64(res), float64(curMax)))
	}

	return res
}

func min(a, b, c int) int {
	if a >= b && c >= b {
		return b
	}
	if a >= c && b >= c {
		return c
	}
	return a
}

func max(a, b, c int) int {
	if a <= b && c <= b {
		return b
	}
	if a <= c && b <= c {
		return c
	}
	return a
}
