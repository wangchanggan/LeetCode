package binary_tree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	var left, right *TreeNode
	if root.Left != nil {
		left = searchBST(root.Left, val)
	}
	if root.Right != nil {
		right = searchBST(root.Right, val)
	}
	if left != nil {
		return left
	}
	return right
}
