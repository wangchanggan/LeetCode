package stack

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	} else if len(heights) == 1 {
		return heights[0]
	}

	var max int
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			var min, count int
			if i == j || i == len(heights)-1 {
				min = heights[i]
				count = 1
			} else {
				min = getMinFromHeights(heights[i : j+1])
				count = j + 1 - i
			}
			if max < min*count {
				max = min * count
			}
		}
	}
	return max
}

func getMinFromHeights(heights []int) int {
	min := heights[0]
	for i := 1; i < len(heights); i++ {
		if min > heights[i] {
			min = heights[i]
		}
	}
	return min
}
