/*Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]


Constraints:

1 <= n <= 8*/

package Medium

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	gen(0, 0, n, "")
	return res
}

func gen(left, right, n int, str string) {
	if left == n && right == n {
		res = append(res, str)
	}
	if left < n {
		gen(left+1, right, n, str+"(")
	}
	if left > right && right < n {
		gen(left, right+1, n, str+")")
	}
}
