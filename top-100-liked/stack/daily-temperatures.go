package stack

func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil {
		return nil
	} else if len(temperatures) == 0 {
		return []int{}
	}

	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return res
}
