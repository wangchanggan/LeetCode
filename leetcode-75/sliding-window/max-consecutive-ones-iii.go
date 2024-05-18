package sliding_window

func longestOnes(nums []int, k int) int {
	/*var max, left, lsum, rsum int
	for right, num := range nums {
		if num == 0 {
			rsum++
		}
		for k < rsum-lsum {
			if nums[left] == 0 {
				lsum++
			}
			left++
		}

		if max < right-left+1 {
			max = right - left + 1
		}
	}
	return max*/
	p := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			p[i+1] = p[i]
		} else {
			p[i+1] = p[i] + 1
		}
	}
	var max int
	for i := 0; i < len(nums); i++ {
		left := binarySearch(p, p[i+1]-k)
		if max < i-left+1 {
			max = i - left + 1
		}
	}

	return max
}

func binarySearch(p []int, target int) int {
	var f, r = 0, len(p) - 1
	for f < r {
		mid := (f + r) / 2
		if p[mid] < target {
			f = mid + 1
		} else {
			r = mid
		}
	}
	return f
}
