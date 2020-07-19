package Easy

func ValidParentheses(s string) bool {
	var array string
	array = array+string('#')
	for i:=0;i<len(s);i++{
		if s[i]=='('||s[i]=='{'||s[i]=='['{
			array = array+string(s[i])
		}else{
			if array[len(array)-1]=='('&&s[i]==')'{
				array = array[:len(array)-1]
			}else if array[len(array)-1]=='{'&&s[i]=='}'{
				array = array[:len(array)-1]
			}else if array[len(array)-1]=='['&&s[i]==']'{
				array = array[:len(array)-1]
			}else{
				array = array+string(s[i])
			}
		}
	}
	if len(array)==1{
		return true
	}
	return false
}