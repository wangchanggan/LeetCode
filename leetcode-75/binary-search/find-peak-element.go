package binary_search

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	f, r := 0, len(nums)-1
	for f < r {
		mid := (f + r) / 2
		if nums[mid] < nums[mid+1] {
			f = mid + 1
		} else {
			r = mid
		}
	}
	return f
}
