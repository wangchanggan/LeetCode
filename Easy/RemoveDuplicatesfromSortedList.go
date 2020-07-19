package Easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveDuplicatesfromSortedList(head *ListNode) *ListNode {
	var result = head
	for result!=nil{
		if result.Next==nil {
			break;
		}
		if result.Val!=result.Next.Val{
			result = result.Next
		}else{
			result.Next = result.Next.Next
		}
	}
	return head
}