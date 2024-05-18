package binary_tree

import "math"

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	treeNodes := make([]*TreeNode, 0)
	treeNodes = append(treeNodes, root)
	var res int
	max := -math.MaxInt
	layer := 0
	for len(treeNodes) > 0 {
		layer++
		var sum int
		l := len(treeNodes)
		for l > 0 {
			treeNode := treeNodes[0]
			treeNodes = treeNodes[1:]
			sum += treeNode.Val
			l--
			if treeNode.Left != nil {
				treeNodes = append(treeNodes, treeNode.Left)
			}
			if treeNode.Right != nil {
				treeNodes = append(treeNodes, treeNode.Right)
			}
		}
		if max < sum {
			max = sum
			res = layer
		}
	}
	return res
}
