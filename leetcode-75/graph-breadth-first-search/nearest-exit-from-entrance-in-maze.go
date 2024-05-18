package graph_breadth_first_search

func nearestExit(maze [][]byte, entrance []int) int {
	if maze == nil || entrance == nil {
		return -1
	}
	vis := make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		vis[i] = make([]int, len(maze[i]))
	}
	roads := make([][2]int, 0)
	roads = append(roads, [2]int{entrance[0], entrance[1]})
	var res int
	for len(roads) > 0 {
		res++
		l := len(roads)
		for l > 0 {
			road := roads[0]
			x := road[0]
			y := road[1]
			roads = roads[1:]
			if x >= 0 && x < len(maze) && y >= 0 && y < len(maze[x]) && vis[x][y] == 0 && maze[x][y] == '.' {
				vis[x][y] = 1
				if (x == 0 || y == 0 || x == len(maze)-1 || y == len(maze[x])-1) && res != 1 {
					return res - 1
				}
				if x+1 >= 0 && x+1 < len(maze) && vis[x+1][y] == 0 && maze[x+1][y] == '.' {
					roads = append(roads, [2]int{x + 1, y})
				}
				if y+1 >= 0 && y+1 < len(maze[x]) && vis[x][y+1] == 0 && maze[x][y+1] == '.' {
					roads = append(roads, [2]int{x, y + 1})
				}
				if x-1 >= 0 && x-1 < len(maze) && vis[x-1][y] == 0 && maze[x-1][y] == '.' {
					roads = append(roads, [2]int{x - 1, y})
				}
				if y-1 >= 0 && y-1 < len(maze[road[0]]) && vis[x][y-1] == 0 && maze[x][y-1] == '.' {
					roads = append(roads, [2]int{x, y - 1})
				}
			}
			l--
		}
	}
	return -1
}
