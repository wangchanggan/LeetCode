package backtrack

import "strconv"

var phoneNumberMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	res := make([]string, 0)
	letterCombinationsOfPhoneNumber(digits, 0, "", &res)
	return res
}

func letterCombinationsOfPhoneNumber(digits string, dIndex int, s string, res *[]string) {
	if dIndex == len(digits) {
		*res = append(*res, s)
		return
	}

	phoneNumber, _ := strconv.Atoi(string(digits[dIndex]))
	phoneNumberStr := phoneNumberMap[phoneNumber]
	for i := 0; i < len(phoneNumberStr); i++ {
		tmp := s + string(phoneNumberStr[i])
		letterCombinationsOfPhoneNumber(digits, dIndex+1, tmp, res)
	}
}
