package Easy

func MaximumDepthofBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaximumDepthofBinaryTree(root.Left)
	right := MaximumDepthofBinaryTree(root.Right)
	if left > right {
		return left + 1
	} else {
		return right+ 1
	}
}