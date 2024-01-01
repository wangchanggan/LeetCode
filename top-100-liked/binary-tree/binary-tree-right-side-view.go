package binary_tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = bfsBinaryTree(root, res)
	return res
}

func bfsBinaryTree(root *TreeNode, res []int) []int {
	treeNodes := make([]*TreeNode, 0)
	treeNodes = append(treeNodes, root)
	for len(treeNodes) > 0 {
		treeNodesLen := len(treeNodes)
		for i := 0; i < treeNodesLen; i++ {
			treeNode := treeNodes[0]
			treeNodes = treeNodes[1:]
			if i == treeNodesLen-1 {
				res = append(res, treeNode.Val)
			}
			if treeNode.Left != nil {
				treeNodes = append(treeNodes, treeNode.Left)
			}
			if treeNode.Right != nil {
				treeNodes = append(treeNodes, treeNode.Right)
			}
		}
	}
	return res
}
