package Easy

func RemoveElement(nums []int, val int) int {
	if len(nums)==0{
		return 0
	}
	var count = -1
	for i:=0;i<len(nums);i++{
		if nums[i]!=val{
			count++
			nums[count] = nums[i]
		}
	}
	return count+1
}