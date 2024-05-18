package binary_search

func minEatingSpeed(piles []int, h int) int {
	f := 0
	var r int
	for _, pile := range piles {
		if pile > r {
			r = pile
		}
	}

	var k int
	t := getTime(piles, r)
	if t <= h {
		k = r
	}

	for f < r {
		speed := (f + r) / 2
		if speed == 0 {
			break
		}
		t := getTime(piles, speed)
		if t <= h {
			k = speed
			r = speed
		} else {
			f = speed + 1
		}
	}

	return k
}

func getTime(piles []int, speed int) int {
	var t int
	for _, pile := range piles {
		var curTime int
		if pile%speed == 0 {
			curTime = pile / speed
		} else {
			curTime = pile/speed + 1
		}
		t += curTime
	}
	return t
}
