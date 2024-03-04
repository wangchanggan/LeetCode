package sliding_window

func longestSubarray(nums []int) int {
	var count, max, f, r int
	var zeroFlag bool
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if !zeroFlag {
				f = count
				zeroFlag = true
			} else {
				r = count
				if max < f+r {
					max = f + r
				}
				f = r
				r = 0
			}
			count = 0
		}
	}

	if r == 0 && max < f+count {
		max = f + count
	}
	if !zeroFlag {
		max--
	}

	return max
}
