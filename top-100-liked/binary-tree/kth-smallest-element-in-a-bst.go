package binary_tree

import "sort"

func kthSmallest(root *TreeNode, k int) int {
	varArr := inorderTraversal(root)
	sort.Ints(varArr)
	return varArr[k-1]
}
