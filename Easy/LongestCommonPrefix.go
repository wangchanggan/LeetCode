package Easy

/**
@brief 返回s中第一个substr实例的索引是否为0
@param[in] s 源字符串
@param[in] substr 目标字符串
@return true为索引为0，false为索引不为0*/
func index(s, substr string) bool{
	var length int
	if len(s)>len(substr){
		length = len(substr)
	}else{
		length = len(s)
	}
	for i:=0;i<length;i++{
		if s[i]!=substr[i]{
			return false
		}
	}
	return true
}

func LongestCommonPrefix(strs []string) string {
	var result string
	if len(strs)==0{
		return ""
	}
	result = strs[0]
	for i:=1;i<len(strs);i++{
		for !index(strs[i], result){
			result = result[0:len(result)-1]
			if result ==""{
				return ""
			}
		}
	}
	for i:=0;i<len(strs);i++{
		if len(strs[i])<len(result){
			result = strs[i]
		}
	}
	return result;
}