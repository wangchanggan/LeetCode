package substring

func subarraySum(nums []int, k int) int {
	if nums == nil {
		return 0
	}
	var res int
	for i := 0; i < len(nums); i++ {
		index := i
		sum := 0
		for index < len(nums) {
			sum += nums[index]
			index++
			if sum == k {
				res++
			}
		}
	}
	return res
}
