package Easy

/*
171. Excel Sheet Column Number
Given a column title as appear in an Excel sheet, return its corresponding column number.
For example:
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
Example 1:
Input: "A"
Output: 1
Example 2:
Input: "AB"
Output: 28
Example 3:
Input: "ZY"
Output: 701
*/

import "math"

func ExcelSheetColumnNumber(s string) int {
	var sum int
	var count int
	for i:=0;i<len(s);i++{
		sum = int(s[i]-64) * int(math.Pow(26, float64(count)))
		count++
	}
	return sum
}