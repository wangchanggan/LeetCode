/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

进阶:
你能不将整数转为字符串来解决这个问题吗？
*/

package Easy

import (
	"math"
	"strconv"
)

func IsPalindrome(x int) bool {
	if x<0{
		return false
	}
	num := x
	var resultNum int
	for x != 0{
		ans := x % 10
		x /= 10
		resultNum = resultNum*10 + ans
		if resultNum > math.MaxInt32 || resultNum < -math.MaxInt32 - 1 {
			return false
		}
	}
	if resultNum != num{
		return false
	}
	return true
}

func IsPalindrome2(x int) bool {
	xStr := strconv.Itoa(x)
	index := 0
	for i := len(xStr)-1; i >= 0; i-- {
		if xStr[i] != xStr[index] {
			return false
		}
		index++
	}
	return true
}