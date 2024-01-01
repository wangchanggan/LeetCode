package double_pointer

func moveZeroes(nums []int) {
	zero := 0
	nonZero := 0
	for zero < len(nums) && nonZero < len(nums) {
		if nums[zero] == 0 && nums[nonZero] != 0 {
			if zero < nonZero {
				nums[zero] = nums[nonZero]
				nums[nonZero] = 0
				zero++
			}
			nonZero++
		} else if nums[zero] == 0 && nums[nonZero] == 0 {
			nonZero++
		} else {
			zero++
		}
	}
}
