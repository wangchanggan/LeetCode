package Easy

func searchTree(l *TreeNode, nums []int){
	if len(nums)==0{
		return
	}
	middle := (len(nums)-1)/2
	l.Val = nums[middle]
	if middle>0{
		l.Left = &TreeNode{0,nil,nil}
		searchTree(l.Left, nums[0:middle])
	}
	if len(nums)>middle+1{
		l.Right = &TreeNode{0,nil,nil}
		searchTree(l.Right, nums[middle+1:])
	}
}

func ConvertSortedArraytoBinarySearchTree(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}else if len(nums)==1{
		return &TreeNode{nums[0],nil,nil}
	}
	middle := (len(nums)-1)/2
	result := &TreeNode{nums[middle],nil,nil}
	if middle>0{
		result.Left = &TreeNode{0,nil,nil}
		searchTree(result.Left, nums[0:middle])
	}
	if len(nums)>middle+1{
		result.Right = &TreeNode{0,nil,nil}
		searchTree(result.Right, nums[middle+1:])
	}
	return result
}