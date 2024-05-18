package Hard

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	h := &ListNode{Val: 0, Next: head}
	pre := h
	cur := h
	var count = 0
	for cur.Next != nil {
		count++
		if count%k == 0 {
			pre, pre.Next = reverseList(pre, cur.Next.Next)
		}
		cur = cur.Next
	}
	return h.Next
}

func reverseList(pre, curNext *ListNode) (*ListNode, *ListNode) {
	p := pre.Next
	var res *ListNode
	for p != curNext {
		res = &ListNode{Val: p.Val, Next: res}
		p = p.Next
	}
	r := res
	for r.Next != nil {
		r = r.Next
	}
	r.Next = curNext
	return r, res
}
