package array_string

import "strings"

func reverseVowels(s string) string {
	if len(s) == 0 {
		return ""
	}

	sBytes := []byte(s)
	f := 0
	r := len(s) - 1
	for f < r {
		for f < len(s) && !strings.Contains("aeiouAEIOU", string(sBytes[f])) {
			f++
		}
		for r > 0 && !strings.Contains("aeiouAEIOU", string(sBytes[r])) {
			r--
		}
		if f < r {
			tmp := sBytes[f]
			sBytes[f] = sBytes[r]
			sBytes[r] = tmp
			f++
			r--
		}
	}
	return string(sBytes)
}
