package Easy

func MinimumDepthofBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		left := MinimumDepthofBinaryTree(root.Left)
		return left + 1
	} else if root.Left == nil && root.Right != nil {
		right := MinimumDepthofBinaryTree(root.Right)
		return right + 1
	}
	left := MinimumDepthofBinaryTree(root.Left)
	right := MinimumDepthofBinaryTree(root.Right)
	if left > right {
		return right + 1
	}
	return left + 1

	// DFS
	/*if root == nil {
		return 0
	}
	result = make([]int, 0)
	minDepthByDFS(root, 1)
	sort.Ints(result)
	return result[0]*/
}

var res []int

func minDepthByDFS(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, level)
		return
	}
	minDepthByDFS(root.Left, level+1)
	minDepthByDFS(root.Right, level+1)
	return
}

func minDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	queue = append(queue, root)
	count := 1
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return count
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count++
	}
	return 0
}
