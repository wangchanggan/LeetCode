package sliding_window

func lengthOfLongestSubstring(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	} else if len(s) == 1 {
		return 1
	} else if len(s) == 2 {
		if s[0] == s[1] {
			return 1
		} else {
			return 2
		}
	}
	for i := 0; i < len(s); i++ {
		sMap := make(map[string]int)
		for j := i; j < len(s); j++ {
			_, ok := sMap[string(s[j])]
			if ok {
				tmp := j - i
				if tmp > res {
					res = tmp
				}
				break
			} else {
				if j == len(s)-1 {
					tmp := j - i
					if tmp > res {
						res = tmp
					}
				}
				sMap[string(s[j])]++
			}
		}
	}
	return res
}
