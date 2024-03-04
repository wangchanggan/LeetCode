package mult_dynamic_programming

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				grid[i][j] = grid[i][j]
			} else if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j]+grid[i][j], grid[i][j-1]+grid[i][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
