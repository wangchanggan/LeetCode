package double_pointer

func maxArea(height []int) int {
	front := 0
	rear := len(height) - 1
	res := 0
	for front < rear {
		area := (rear - front) * min(height[front], height[rear])
		res = max(res, area)
		if height[front] < height[rear] {
			front++
		} else {
			rear--
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
