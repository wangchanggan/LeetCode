package Easy

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/
func TwoSum(nums []int, target int) []int {
	var results []int
	for i:=0;i<len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			sum := nums[i] + nums[j]
			if sum == target{
				results = append(results, i)
				results = append(results, j)
			}
		}
	}
	return results
}