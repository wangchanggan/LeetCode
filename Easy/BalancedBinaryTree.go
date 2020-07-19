package Easy

func depthofBinaryTree(root *TreeNode) (int,bool){
	if root == nil {
		return 0,true
	}
	left,flag1 := depthofBinaryTree(root.Left)
	right,flag2 := depthofBinaryTree(root.Right)
	if !flag1 || !flag2 {
		return 0, false
	}
	if left-right>1 ||right-left>1{
		return 0, false
	}
	if left>right {
		return left+1, true
	}
	return right+1, true
}


func BalancedBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	flag := true
	_,flag = depthofBinaryTree(root)
	return flag
}