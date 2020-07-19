package Easy

func SingleNumber(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	result := nums[0]
	for i:=1;i<len(nums);i++{
		flag := true
		for j:=0;j<i;j++{
			if nums[j]==nums[i]{
				flag = false
				break
			}
		}
		for k:=i+1;k<len(nums);k++{
			if nums[k]==nums[i]{
				flag = false
				break
			}
		}
		if flag{
			result = nums[i]
		}
	}
	return result
}