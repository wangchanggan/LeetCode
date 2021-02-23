package Easy

func MaximumDepthofBinaryTree(root *TreeNode) int {
	/*if root == nil {
		return 0
	}
	left := MaximumDepthofBinaryTree(root.Left)
	right := MaximumDepthofBinaryTree(root.Right)
	if left > right {
		return left + 1
	} else {
		return right+ 1
	}*/

	/*result = make([][]int, 0)
	maxDepthByDFS(root, 0)
	return len(result)*/

	return len(maxDepthByBFS(root))
}

var result [][]int

func maxDepthByDFS(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(result) == level {
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	maxDepthByDFS(root.Left, level+1)
	maxDepthByDFS(root.Right, level+1)
}

func maxDepthByBFS(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		list := make([]int, 0)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)
	}
	return res
}
