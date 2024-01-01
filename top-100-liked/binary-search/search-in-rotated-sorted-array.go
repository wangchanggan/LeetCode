package binary_search

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[i] <= nums[mid] {
			// i到mid是有序的
			if nums[i] <= target && target < nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else {
			// i到mid是无序的，那mid到j就是有序的
			if nums[mid] < target && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}
	return -1
}
