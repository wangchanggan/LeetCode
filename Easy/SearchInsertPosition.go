package Easy

func SearchInsertPosition(nums []int, target int) int {
	for i:=0;i<len(nums);i++{
		if nums[i]==target{
			return i
		}
	}
	for i:=0;i<len(nums);i++{
		if nums[i]>target{
			return i-1
		}
	}
	return len(nums)
}
