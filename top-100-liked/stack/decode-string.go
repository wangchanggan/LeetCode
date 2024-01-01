package stack

import (
	"strconv"
)

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var res string
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			var numIndex, num int
			var str string
			// get string
			l := len(stack)
			for j := l - 1; j >= 0; j-- {
				if stack[j] == '[' {
					str = string(stack[j+1:])
					stack = stack[:j]
					numIndex = j
					break
				}
			}
			// get num
			l = len(stack)
			j := 0
			for j = l - 1; j >= 0; j-- {
				if '0' <= stack[j] && stack[j] <= '9' {
					continue
				} else {
					break
				}
			}
			num = stringToInt(string(stack[j+1 : numIndex]))
			stack = stack[:j+1]
			newStr := generateDecodeString(str, num)
			for j := 0; j < len(newStr); j++ {
				stack = append(stack, newStr[j])
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	for _, v := range stack {
		res += string(v)
	}
	return res
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func generateDecodeString(s string, num int) string {
	var str string
	for i := 0; i < num; i++ {
		str += s
	}
	return str
}
