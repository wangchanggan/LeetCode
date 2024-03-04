package array_string

func gcdOfStrings(str1 string, str2 string) string {
	for i := len(str2); i > 0; i-- {
		if len(str1)%i == 0 && len(str2)%i == 0 {
			for j := 0; j <= len(str2)-i; j++ {
				substr := str2[j : j+i]
				if greatestCommonDivisorOfStrings(str1, str2, substr) {
					return substr
				}
			}
		}
	}
	return ""
}

func greatestCommonDivisorOfStrings(str1, str2, substr string) bool {
	if len(str1)%len(substr) == 0 {
		var res string
		for i := 0; i < len(str1)/len(substr); i++ {
			res += substr
		}
		if res != str1 {
			return false
		}
	} else {
		return false
	}
	if len(str2)%len(substr) == 0 {
		var res string
		for i := 0; i < len(str2)/len(substr); i++ {
			res += substr
		}
		if res != str2 {
			return false
		}
	} else {
		return false
	}
	return true
}
