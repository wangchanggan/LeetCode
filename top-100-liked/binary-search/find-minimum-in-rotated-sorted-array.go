package binary_search

func findMin(nums []int) int {
	if nums == nil {
		return -1
	} else if len(nums) == 1 {
		return nums[0]
	}
	min := nums[0]
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[i] <= nums[mid] {
			// i到mid是有序的
			if min > nums[i] {
				min = nums[i]
			}
			i = mid + 1
		} else {
			// mid到j是有序的
			if min > nums[mid] {
				min = nums[mid]
			}
			j = mid - 1
		}
	}
	return min
}
