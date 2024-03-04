package prefix_summary

func largestAltitude(gain []int) int {
	var max, sum int
	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		if max < sum {
			max = sum
		}
	}
	return max
}
