package linked_list

import "strconv"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lNumStr := numStrAdd(getNumStrFromListNode(l1), getNumStrFromListNode(l2))
	l := new(ListNode)
	h := l
	for i := len(lNumStr) - 1; i >= 0; i-- {
		h.Val, _ = strconv.Atoi(string(lNumStr[i]))
		if i != 0 {
			h.Next = new(ListNode)
			h = h.Next
		}
	}
	return l
}

func getNumStrFromListNode(l *ListNode) string {
	str := ""
	for l != nil {
		str = strconv.Itoa(l.Val) + str
		l = l.Next
	}
	return str
}

func numStrAdd(ns1, ns2 string) string {
	var shortStr, longStr string
	if len(ns1) > len(ns2) {
		shortStr = ns2
		longStr = ns1
	} else {
		shortStr = ns1
		longStr = ns2
	}

	var res string
	var count int
	longStrIndex := len(longStr) - 1
	for i := len(shortStr) - 1; i >= 0; i-- {
		shortNum, _ := strconv.Atoi(string(shortStr[i]))
		longNum, _ := strconv.Atoi(string(longStr[longStrIndex]))
		sumNum := shortNum + longNum + count
		if sumNum >= 10 {
			count = sumNum / 10
			res = strconv.Itoa(sumNum%10) + res
		} else {
			count = 0
			res = strconv.Itoa(sumNum) + res
		}
		longStrIndex--
	}

	for i := len(longStr) - len(shortStr) - 1; i >= 0; i-- {
		longNum, _ := strconv.Atoi(string(longStr[i]))
		sumNum := longNum + count
		if sumNum >= 10 {
			count = sumNum / 10
			res = strconv.Itoa(sumNum%10) + res
		} else {
			count = 0
			res = strconv.Itoa(sumNum) + res
		}
	}

	if count != 0 {
		res = strconv.Itoa(count) + res
	}
	return res
}
