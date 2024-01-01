package binary_tree

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var total int
	res := new([]int)
	pathSumBinaryTree(root, res)
	for i := 0; i < len(*res); i++ {
		if (*res)[i] == targetSum {
			total++
		}
	}
	return total
}

func pathSumBinaryTree(root *TreeNode, res *[]int) {
	dfsBinaryTree(root, 0, res)
	if root != nil {
		pathSumBinaryTree(root.Left, res)
		pathSumBinaryTree(root.Right, res)
	}
}

func dfsBinaryTree(root *TreeNode, sum int, res *[]int) {
	if root == nil {
		return
	}
	sum += root.Val
	*res = append(*res, sum)
	dfsBinaryTree(root.Left, sum, res)
	dfsBinaryTree(root.Right, sum, res)
}
