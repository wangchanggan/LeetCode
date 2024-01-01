package regular_array

func firstMissingPositive(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	var max int
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
		if max < nums[i] {
			max = nums[i]
		}
	}

	for i := 1; i <= max; i++ {
		if _, ok := numsMap[i]; !ok {
			return i
		}
	}

	if max >= 0 {
		return max + 1
	}
	return -1
}
