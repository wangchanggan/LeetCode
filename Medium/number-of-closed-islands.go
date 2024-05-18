package Medium

func closedIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	vis := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		vis[i] = make([]int, len(grid[i]))
	}
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 && vis[i][j] == 0 {
				edge := new(bool)
				dfsIsland(grid, vis, i, j, edge)
				if !*edge {
					res++
				}
			}
		}
	}
	return res
}

func dfsIsland(grid, vis [][]int, i, j int, edge *bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || vis[i][j] == 1 || grid[i][j] == 1 {
		return
	}
	vis[i][j] = 1
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
		*edge = true
	}
	dfsIsland(grid, vis, i+1, j, edge)
	dfsIsland(grid, vis, i, j+1, edge)
	dfsIsland(grid, vis, i-1, j, edge)
	dfsIsland(grid, vis, i, j-1, edge)
}
