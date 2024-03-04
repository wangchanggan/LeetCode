package array_string

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if candies == nil || len(candies) == 0 {
		return nil
	}

	var max int
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}
