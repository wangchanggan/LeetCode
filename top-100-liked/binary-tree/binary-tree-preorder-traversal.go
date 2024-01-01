package binary_tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = binaryTreePreorderTraversal(root, res)
	return res
}

func binaryTreePreorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	if root.Left != nil {
		res = binaryTreePreorderTraversal(root.Left, res)
	}
	if root.Right != nil {
		res = binaryTreePreorderTraversal(root.Right, res)
	}
	return res
}
