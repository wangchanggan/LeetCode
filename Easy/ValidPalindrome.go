package Easy

import (
	"strings"
)

func ValidPalindrome(s string) bool {
	var str string
	s = strings.ToUpper(s)
	for i:=0;i<len(s);i++{
		if s[i]>='a'&&s[i]<='z'{
			str += string(s[i])
		}
		if s[i]>='A'&&s[i]<='Z'{
			str += string(s[i])
		}
		if s[i]>='0'&&s[i]<='9'{
			str += string(s[i])
		}
	}
	l := len(str)-1
	for i:=0;i<l;i++{
		if str[i]!=str[l]{
			return false
		}
		l--
	}
	return true
}