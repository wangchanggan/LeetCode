package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := head
	var length int
	for l != nil {
		length++
		l = l.Next
	}

	if length < n {
		return nil
	} else if length == n {
		return head.Next
	}

	l = head
	for i := 0; i < length-n-1; i++ {
		l = l.Next
	}
	if l.Next != nil {
		if l.Next.Next != nil {
			l.Next = l.Next.Next
		} else {
			l.Next = nil
		}
	} else {
		return nil
	}

	return head
}
