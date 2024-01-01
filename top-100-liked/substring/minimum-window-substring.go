package substring

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	} else if s == t {
		return s
	}

	var res string
	tCnt := getLetterMap(t)
	for i := 0; i <= len(s)-len(t); i++ {
		for j := i + len(t); j <= len(s); j++ {
			sCnt := getLetterMap(s[i:j])
			flag := true
			for tk, tv := range tCnt {
				sv, ok := sCnt[tk]
				if !ok || tv > sv {
					flag = false
					break
				}
			}
			if flag {
				if res == "" || len(s[i:j]) < len(res) {
					res = s[i:j]
				}
				break
			}
		}
	}
	return res
}

func getLetterMap(s string) map[string]int {
	cnt := make(map[string]int)
	for i := 0; i < len(s); i++ {
		cnt[string(s[i])]++
	}
	return cnt
}
