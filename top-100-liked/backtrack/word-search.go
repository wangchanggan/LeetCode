package backtrack

func exist(board [][]byte, word string) bool {
	if board == nil || len(word) == 0 {
		return false
	}

	vis := make([][]int, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]int, len(board[i]))
	}

	w := word[0]
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == w {
				var res bool
				wordSearch(board, deepCopySlices(vis), i, j, 0, word, "", &res)
				if res {
					return true
				}
			}
		}
	}
	return false
}

func wordSearch(board [][]byte, vis [][]int, x, y, index int, word, subWord string, res *bool) {
	if 0 > x || x >= len(board) || 0 > y || y >= len(board[x]) || vis[x][y] == 1 || *res {
		return
	}
	if word[index] == board[x][y] {
		subWord += string(board[x][y])
	} else {
		return
	}

	vis[x][y] = 1
	if word == subWord {
		*res = true
		return
	}
	wordSearch(board, deepCopySlices(vis), x+1, y, index+1, word, subWord, res)
	wordSearch(board, deepCopySlices(vis), x-1, y, index+1, word, subWord, res)
	wordSearch(board, deepCopySlices(vis), x, y+1, index+1, word, subWord, res)
	wordSearch(board, deepCopySlices(vis), x, y-1, index+1, word, subWord, res)
}

func deepCopySlices(a [][]int) [][]int {
	b := make([][]int, len(a))
	for i := 0; i < len(b); i++ {
		b[i] = make([]int, len(a[i]))
		for j := 0; j < len(a[i]); j++ {
			b[i][j] = a[i][j]
		}
	}
	return b
}
