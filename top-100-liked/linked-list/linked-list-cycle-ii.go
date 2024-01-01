package linked_list

func detectCycle(head *ListNode) *ListNode {
	headMap := make(map[*ListNode]int)
	for head != nil {
		if _, ok := headMap[head]; ok {
			return head
		} else {
			headMap[head]++
		}
		head = head.Next
	}
	return nil
}
