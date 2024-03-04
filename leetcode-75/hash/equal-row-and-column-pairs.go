package hash

func equalPairs(grid [][]int) int {
	arr := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		arr = append(arr, grid[i])
	}
	for i := 0; i < len(grid); i++ {
		tmp := make([]int, 0)
		for j := 0; j < len(grid); j++ {
			tmp = append(tmp, grid[j][i])
		}
		arr = append(arr, tmp)
	}
	var res int
	mid := len(arr) / 2
	for i := 0; i < mid; i++ {
		for j := mid; j < len(arr); j++ {
			if equalArrays(arr[i], arr[j]) {
				res++
			}
		}
	}
	return res
}

func equalArrays(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
