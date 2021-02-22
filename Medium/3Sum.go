/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []


Constraints:

0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

package Medium

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	result := make([][]int, 0)
	type KeyArray [3]int
	tripletsMap := make(map[KeyArray][]int)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				tmp := []int{nums[i], nums[left], nums[right]}
				_, ok := tripletsMap[[3]int{nums[i], nums[left], nums[right]}]
				if !ok {
					tripletsMap[[3]int{nums[i], nums[left], nums[right]}] = tmp
					result = append(result, tmp)
				}

				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}

	}

	return result
}
