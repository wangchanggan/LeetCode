package Easy

func strStr(l1, l2, length1,length2 int,haystack string, needle string) bool{
	for l1<length1&&l2<length2{
		l1++
		l2++
		if haystack[l1]!=needle[l2]{
			return false
		}
	}
	return true
}

func ImplementstrStr(haystack string, needle string) int {
	if len(needle)==0{
		return 0
	}else if len(haystack)<len(needle){
		return -1
	}
	for i:=0;i<len(haystack);i++{
		if haystack[i]==needle[0]{
			if strStr(i,0,len(haystack)-1,len(needle)-1,haystack,needle)&&len(haystack)-i>len(needle){
				return i
			}
		}
	}
	return -1
}