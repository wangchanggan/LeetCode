package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	headArr := make([]int, 0)
	for head != nil {
		headArr = append(headArr, head.Val)
		head = head.Next
	}

	rear := len(headArr) - 1
	for i := 0; i < len(headArr); i++ {
		if headArr[i] != headArr[rear] {
			return false
		}
		rear--
	}
	return true
}
