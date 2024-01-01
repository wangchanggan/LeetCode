package binary_tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
