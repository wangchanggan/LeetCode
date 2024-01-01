package dynamic_programming

import "strings"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}

	newWordDict := make([]string, 0)
	for i := 0; i < len(wordDict); i++ {
		if strings.Contains(s, wordDict[i]) {
			newWordDict = append(newWordDict, wordDict[i])
		}
	}

	if len(newWordDict) == 0 {
		return false
	}

	return false
}
