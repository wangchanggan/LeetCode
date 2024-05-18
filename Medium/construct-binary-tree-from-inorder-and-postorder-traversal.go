package Medium

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 || postorder == nil || len(postorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i+1:len(postorder)-1])

	return root
}
