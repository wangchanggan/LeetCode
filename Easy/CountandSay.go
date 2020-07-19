package Easy

import (
	"strconv"
)

func CountandSay(n int) string {
	if n==1{
		return "1"
	}else if n==2{
		return "11"
	}else if n==3{
		return "21"
	}
	countSay := "21"
	for i:=4;i<=n;i++{
		var countNum string
		var nums int
		count := 0
		for j:=1;j<len(countSay);j++{
			if countSay[j-1]!=countSay[j]{
				num, _ := strconv.Atoi(string(countSay[j-1]))
				countNum = countNum + strconv.Itoa(j-count)+strconv.Itoa(num)
				count = j
			}
		}
		num, _ := strconv.Atoi(string(countSay[len(countSay)-1]))
		nums = len(countSay)-count
		countSay = countNum + strconv.Itoa(nums) + strconv.Itoa(num)
	}
	return countSay
}