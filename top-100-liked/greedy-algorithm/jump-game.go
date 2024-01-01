package greedy_algorithm

func CanJump(nums []int) bool {
	jumpKey := 0
	jumpValue := nums[jumpKey]
	for jumpKey < len(nums) {
		jumpKey += jumpValue
		if jumpKey >= len(nums)-1 {
			break
		}
		jumpValue = nums[jumpKey]
		if jumpValue == 0 {
			return false
		}
	}
	if jumpKey >= len(nums)-1 {
		return true
	}
	return false
}
