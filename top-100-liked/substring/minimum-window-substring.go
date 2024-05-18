package substring

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	} else if s == t {
		return s
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	ansL, ansR := -1, -1
	l, r := 0, 0
	ansLen := math.MaxInt32
	sMap := make(map[byte]int)
	var res string
	for r < len(s) {
		if tMap[s[r]] > 0 {
			sMap[s[r]]++
		}

		for l <= r {
			flag := true
			for k, v := range tMap {
				if sMap[k] < v {
					flag = false
					break
				}
			}
			if flag {
				if r-l+1 < ansLen {
					ansLen = r - l + 1
					ansL, ansR = l, l+ansLen
					res = s[ansL:ansR]
				}
				if _, ok := tMap[s[l]]; ok {
					sMap[s[l]]--
				}
				l++
			} else {
				break
			}
		}
		r++
	}
	return res
}
