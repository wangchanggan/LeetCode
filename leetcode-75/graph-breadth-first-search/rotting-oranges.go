package graph_breadth_first_search

func orangesRotting(grid [][]int) int {
	if grid == nil {
		return -1
	}

	oranges := make([][2]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				oranges = append(oranges, [2]int{i, j})
			}
		}
	}
	vis := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		vis[i] = make([]int, len(grid[i]))
	}

	var res, x, y int
	for len(oranges) > 0 {
		l := len(oranges)
		res++
		for l > 0 {
			orange := oranges[0]
			oranges = oranges[1:]
			x = orange[0]
			y = orange[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && (grid[x][y] == 1 || grid[x][y] == 2) {
				grid[x][y] = 2
				vis[x][y] = 1
				if x+1 >= 0 && x+1 < len(grid) && vis[x+1][y] == 0 && grid[x+1][y] == 1 {
					vis[x+1][y] = 1
					oranges = append(oranges, [2]int{x + 1, y})
				}
				if y+1 >= 0 && y+1 < len(grid[x]) && vis[x][y+1] == 0 && grid[x][y+1] == 1 {
					vis[x][y+1] = 1
					oranges = append(oranges, [2]int{x, y + 1})
				}
				if x-1 >= 0 && x-1 < len(grid) && vis[x-1][y] == 0 && grid[x-1][y] == 1 {
					vis[x-1][y] = 1
					oranges = append(oranges, [2]int{x - 1, y})
				}
				if y-1 >= 0 && y-1 < len(grid[x]) && vis[x][y-1] == 0 && grid[x][y-1] == 1 {
					vis[x][y-1] = 1
					oranges = append(oranges, [2]int{x, y - 1})
				}
			}
			l--
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if res == 0 {
		return 0
	}
	return res - 1
}
