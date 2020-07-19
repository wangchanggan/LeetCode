package Easy

func minCha(a,b int) (int,int){
	if a>b{
		return a-b,0
	}
	return b-a,1
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var lengthA,lengthB int
	a,b := headA, headB
	for a!=nil{
		lengthA++
		a = a.Next
	}
	for b!=nil{
		lengthB++
		b = b.Next
	}

	if (lengthA < lengthB) {
		for i:=0;i<lengthB-lengthA;i++{
			headB = headB.Next
		}
	} else {
		for i:=0;i<lengthA-lengthB;i++{
			headA = headA.Next
		}
	}
	for headA != nil && headB != nil  && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	if headA !=nil {
		return headA
	}else if headB !=nil {
		return headB
	}
	return nil
}