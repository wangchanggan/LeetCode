package double_pointer

func trap(height []int) int {
	index := 0
	res := 0
	for index < len(height)-1 {
		var flag bool
		for i := index + 1; i < len(height); i++ {
			if height[index] != 0 && height[i] != 0 && height[index] <= height[i] {
				res += trapping(height[index : i+1])
				index = i
				flag = true
				break
			}
		}
		if !flag {
			index++
		}
	}
	return res
}

func trapping(height []int) int {
	l := len(height)
	total := min(height[0], height[l-1]) * (len(height) - 2)
	for i := 0; i < l; i++ {
		if i != 0 && i != l-1 {
			total = total - height[i]
		}
	}
	return total
}
