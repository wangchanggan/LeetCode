package Easy

func MaximumSubarray(nums []int) int {
	if len(nums)==0{
		return 0
	}
	var count int
	max := -2147483647
	for i:=0;i<len(nums);i++{
		if count+nums[i]<nums[i] {
			count = nums[i]
		} else {
			count += nums[i]
		}
		if count > max {
			max = count
		}
	}
	return max

}