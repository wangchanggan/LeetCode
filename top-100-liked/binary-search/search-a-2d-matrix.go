package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] == target || matrix[i][n-1] == target {
			return true
		}
		if matrix[i][0] < target && target < matrix[i][n-1] {
			return searchArray(matrix[i], target, 1, n-2)
		}
	}
	return false
}

func searchArray(arr []int, target, i, j int) bool {
	if i > j || i < 0 || j < 0 || i >= len(arr) || j >= len(arr) {
		return false
	}
	mid := (i + j) / 2
	if arr[mid] == target {
		return true
	} else if arr[mid] < target {
		return searchArray(arr, target, mid+1, j)
	} else {
		return searchArray(arr, target, i, mid-1)
	}
}
