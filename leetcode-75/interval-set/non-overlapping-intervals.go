package interval_set

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	interval := intervals[0][1]
	res := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= interval {
			res++
			interval = intervals[i][1]
		}
	}
	return len(intervals) - res
}
