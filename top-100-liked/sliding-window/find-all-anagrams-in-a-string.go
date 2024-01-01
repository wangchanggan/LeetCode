package sliding_window

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	res := make([]int, 0)
	for i := 0; i <= len(s)-len(p); i++ {
		ss := s[i : i+len(p)]
		if isAnagram(ss, p) {
			res = append(res, i)
		}
	}
	return res
}

func isAnagram(s1, s2 string) bool {
	s1Cnt := [26]int{}
	for _, s := range s1 {
		s1Cnt[s-'a']++
	}
	s2Cnt := [26]int{}
	for _, s := range s2 {
		s2Cnt[s-'a']++
	}
	for i := 0; i < len(s1Cnt); i++ {
		if s1Cnt[i] != s2Cnt[i] {
			return false
		}
	}
	return true
}
