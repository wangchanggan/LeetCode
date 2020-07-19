package Easy

import "fmt"

//import "fmt"

/*
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

1 -> A
2 -> B
3 -> C
...
26 -> Z
27 -> AA
28 -> AB
...

*/

func ConvertToTitle(n int) string {
	var result string
	var answer string
	for n > 26{
		char := n % 26
		n = n / 26
		if char == 0{
			char = 26
			n -= 1
		}
		result = result + string(char+64)
		fmt.Println(result)
	}
	result = result + string(n+64)
	for i:=len(result)-1;i>=0;i--{
		answer = answer + string(result[i])
	}
	return answer
}