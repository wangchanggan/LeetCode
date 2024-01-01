package binary_tree

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	treeNodeStack := make([]*TreeNode, 0)
	treeNodeStack = append(treeNodeStack, root)
	for len(treeNodeStack) > 0 {
		treeNode := treeNodeStack[len(treeNodeStack)-1]
		treeNodeStack = treeNodeStack[:len(treeNodeStack)-1]
		if treeNode.Left != nil {
			tmp := new(TreeNode)
			tmp.Val = treeNode.Val
			if treeNode.Right != nil {
				tmp.Right = treeNode.Right
			}
			treeNodeStack = append(treeNodeStack, tmp)
			treeNodeStack = append(treeNodeStack, treeNode.Left)
		} else {
			if treeNode.Right != nil {
				tmp := new(TreeNode)
				tmp.Val = treeNode.Val
				treeNodeStack = append(treeNodeStack, tmp)
				treeNodeStack = append(treeNodeStack, treeNode.Right)
			} else {
				res = append(res, treeNode.Val)
			}
		}
	}
	return res
}
