package backtrack

func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	res := make([][]string, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, []string{string(s[i])})
		for j := i; j < len(s); j++ {
			if i != j && i != len(s)-1 && isPalindrome(s[i:j]) {
				res = append(res, []string{s[i:j]})
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	front := 0
	rear := len(s) - 1
	for front < rear {
		if s[front] != s[rear] {
			return false
		}
	}
	return true
}
