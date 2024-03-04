package sliding_window

func maxVowels(s string, k int) int {
	var count, max int
	for i := 0; i < k; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			count++
		}
	}
	if max < count {
		max = count
	}
	for i := k; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			count++
		}
		if s[i-k] == 'a' || s[i-k] == 'e' || s[i-k] == 'i' || s[i-k] == 'o' || s[i-k] == 'u' {
			count--
		}
		if max < count {
			max = count
		}
	}
	return max
}
