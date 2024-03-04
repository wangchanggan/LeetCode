package array_string

func reverseWords(s string) string {
	ss := make([]string, 0)
	f := -1
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if f == -1 {
				f = i
			}
		} else {
			if f != -1 {
				ss = append(ss, s[f:i])
				f = -1
			}
		}
		if i == len(s)-1 && f != -1 {
			ss = append(ss, s[f:len(s)])
		}
	}
	var res string
	for i := len(ss) - 1; i >= 0; i-- {
		if i != 0 {
			res = res + ss[i] + " "
		} else {
			res += ss[i]
		}
	}
	return res
}
