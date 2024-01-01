package regular_array

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	r := make([]int, len(nums))
	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = l[i] * r[i]
	}
	return res
}
