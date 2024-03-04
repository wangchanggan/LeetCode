package sliding_window

import "fmt"

func longestOnes(nums []int, k int) int {
	var count, max, threeZero int
	zeroCount := k
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		if nums[i] == 1 {
			count++
			threeZero = 0
		} else {
			threeZero++
			if zeroCount > 0 {
				zeroCount--
				count++
			} else {
				if threeZero == 3 {
					zeroCount = k
					count = 0
					threeZero = 0
					i--
				}
			}
		}
		if max < count {
			max = count
		}
	}
	return max
}
