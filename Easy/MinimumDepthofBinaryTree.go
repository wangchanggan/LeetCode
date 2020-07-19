package Easy

func MinimumDepthofBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		left := MinimumDepthofBinaryTree(root.Left)
		return left + 1
	} else if root.Left == nil && root.Right != nil {
		right := MinimumDepthofBinaryTree(root.Right)
		return right + 1
	}
	left := MinimumDepthofBinaryTree(root.Left)
	right := MinimumDepthofBinaryTree(root.Right)
	if left>right{
		return right+1
	}
	return left+1
}