package stack

func removeStars(s string) string {
	if len(s) == 0 {
		return ""
	} else if len(s) == 1 {
		if s[0] == '*' {
			return ""
		} else {
			return s
		}
	}

	res := string(s[0])
	for i := 1; i < len(s); i++ {
		if s[i] == '*' {
			res = res[:len(res)-1]
		} else {
			res += string(s[i])
		}
	}
	return res
}
