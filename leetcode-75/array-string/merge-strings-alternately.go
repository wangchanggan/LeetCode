package array_string

func mergeAlternately(word1 string, word2 string) string {
	var res string
	for i := 0; i < len(word1); i++ {
		res += string(word1[i])
		if i < len(word2) {
			res += string(word2[i])
		}
	}

	if len(word1) < len(word2) {
		for i := len(word1); i < len(word2); i++ {
			res += string(word2[i])
		}
	}
	return res
}
