package Easy

type ListNode struct {
	Val int
	Next *ListNode
}
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = &ListNode{0,nil}
	var result = l
	for l1!=nil && l2!=nil{
		if l1.Val>l2.Val{
			l = l2
			l2 = l2.Next
			l = l.Next
		}else{
			l = l1
			l1 = l1.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l = l1
	}else if l2 != nil {
		l = l2
	}
	return result
}