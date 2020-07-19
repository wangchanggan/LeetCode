package Easy

func RemoveDuplicatesfromSortedArray(nums []int) int {
	if len(nums)==0{
		return 0
	}
	var count int
	var num int
	for i:=0;i<len(nums)-1;i++{
		if nums[i]!=nums[i+1]{
			nums[count+1] = nums[i+1]
			num++
			count++
		}
	}
	return num+1
}