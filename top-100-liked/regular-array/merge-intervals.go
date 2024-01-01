package regular_array

import "sort"

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][0] == intervals[i][0] {
			if res[len(res)-1][1] < intervals[i][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			if res[len(res)-1][1] < intervals[i][0] {
				res = append(res, intervals[i])
			} else {
				if res[len(res)-1][1] < intervals[i][1] {
					res[len(res)-1][1] = intervals[i][1]
				}
			}
		}
	}
	return res
}
