package binary_tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	r := root
	invertBinaryTree(r)
	return root
}

func invertBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	if root.Left != nil {
		invertBinaryTree(root.Left)
	}
	if root.Right != nil {
		invertBinaryTree(root.Right)
	}
}
