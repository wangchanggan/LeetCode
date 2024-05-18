package binary_search

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	if spells == nil || potions == nil {
		return nil
	}

	sort.Ints(potions)
	res := make([]int, 0)
	for i := 0; i < len(spells); i++ {
		var target int64
		if success%int64(spells[i]) == 0 {
			target = success / int64(spells[i])
		} else {
			target = success/int64(spells[i]) + 1
		}

		f, r := 0, len(potions)-1
		var mid int
		for f < r {
			mid = (f + r) / 2
			if int64(potions[mid]) < target {
				f = mid + 1
			} else {
				r = mid
			}
		}
		if f == len(potions)-1 && int64(potions[f]) < target {
			res = append(res, 0)
		} else if f == 0 && int64(potions[f]) >= target {
			res = append(res, len(potions))
		} else {
			res = append(res, len(potions)-f)
		}
	}
	return res
}
