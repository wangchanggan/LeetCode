package Easy

func LengthOfLastWord(s string) int {
	if len(s)==0{
		return 0
	}
	var result string
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			result = s[i:]
			break
		}
	}
	for i:=len(result)-1;i>=0;i--{
		if s[i]!=' '{
			result = s[:i+1]
			break
		}
	}
	for i:=len(result)-1;i>=0;i--{
		if result[i]==' '{
			return len(result)-i-1
		}
	}
	return len(result)
}