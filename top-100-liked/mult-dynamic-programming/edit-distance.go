package mult_dynamic_programming

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	} else if len(word1) == 0 && len(word2) != 0 {
		return len(word2)
	} else if len(word1) != 0 && len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				if i == 0 && j == 0 {
					dp[i][j] = 0
				} else if i == 0 && j != 0 {
					dp[i][j] = j
				} else if i != 0 && j == 0 {
					dp[i][j] = i
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				} else if i == 0 && j != 0 {
					dp[i][j] = dp[i][j-1] + 1
				} else if i != 0 && j == 0 {
					dp[i][j] = dp[i-1][j] + 1
				} else {
					dp[i][j] = minThree(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
			}
		}
	}
	return dp[len(word1)-1][len(word2)-1]
}

func minThree(a, b, c int) int {
	if a < b {
		if c < a {
			return c
		} else {
			return a
		}
	} else {
		if b > c {
			return c
		} else {
			return b
		}
	}
}
