package Easy

import (
	"strconv"
)

func AddBinary(a string, b string) string {
	var count int
	var result string
	i := len(a)-1
	j := len(b)-1
	for i>=0&&j>=0{
		num1,_ := strconv.Atoi(string(a[i]))
		num2,_ := strconv.Atoi(string(b[j]))
		if num1==1&&num2==1&&count==1{
			count = 1
			result += "1"
		}else if num1==1&&num2==0&&count==0{
			count = 0
			result += "1"
		}else if num1==0&&num2==1&&count==0{
			count = 0
			result += "1"
		}else if num1==1&&num2==0&&count==1{
			count = 1
			result += "0"
		}else if num1==0&&num2==1&&count==1{
			count = 1
			result += "0"
		}else if num1==0&&num2==0&&count==0{
			count = 0
			result += "0"
		}else if num1==1&&num2==1&&count==0{
			count = 1
			result += "0"
		}else if num1==0&&num2==0&&count==1{
			count = 0
			result += "1"
		}else{
			count = 0
			result += "0"
		}
		i--
		j--
	}
	if i>=0{
		for k:=i;k>=0;k--{
			num1,_ := strconv.Atoi(string(a[k]))
			if num1==1&&count==1{
				count = 1
				result += "0"
			}else if num1==0&&count==0{
				count = 0
				result += "0"
			}else{
				count = 0
				result += "1"
			}
		}
	}else if j>=0{
		for k:=j;k>=0;k--{
			num2,_ := strconv.Atoi(string(b[k]))
			if num2+count>=2{
				count = 1
				result += "0"
			}else if num2==0&&count==0{
				count = 0
				result += "0"
			}else{
				count = 0
				result += "1"
			}
		}
	}
	if count==1{
		result += "1"
	}
	var answer string
	for i:=len(result)-1;i>=0;i--{
		answer += string(result[i])
	}
	return answer
}