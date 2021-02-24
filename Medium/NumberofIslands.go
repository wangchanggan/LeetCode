/*Given an m x n 2d grid map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input: grid = [
["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
["1","1","0","0","0"],
["1","1","0","0","0"],
["0","0","1","0","0"],
["0","0","0","1","1"]
]
Output: 3


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.*/

package Medium

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				floodFill(grid, i, j)
			}
		}
	}
	return count
}

func floodFill(gird [][]byte, i, j int) {
	if i >= 0 && i < len(gird) && j >= 0 && j < len(gird[i]) {
		if gird[i][j] == '1' {
			gird[i][j] = '0'
			floodFill(gird, i+1, j)
			floodFill(gird, i, j+1)
			floodFill(gird, i-1, j)
			floodFill(gird, i, j-1)
		}
	}
}
