package hash

import "sort"

func uniqueOccurrences(arr []int) bool {
	arrMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		arrMap[arr[i]]++
	}

	a := make([]int, 0)
	i := 0
	for _, v := range arrMap {
		a = append(a, v)
		i++
	}
	sort.Ints(a)
	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			return false
		}
	}
	return true
}
