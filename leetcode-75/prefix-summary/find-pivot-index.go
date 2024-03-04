package prefix_summary

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if sumNums(nums[1:]) == 0 {
				return i
			}
		} else if i == len(nums)-1 {
			if sumNums(nums[:len(nums)-1]) == 0 {
				return i
			}
		} else {
			if sumNums(nums[:i]) == sumNums(nums[i+1:]) {
				return i
			}
		}
	}
	return -1
}

func sumNums(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
