/*Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000*/

package Medium

func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	levelOrderByDFS(root, 0)
	return result
}

var result [][]int

func levelOrderByDFS(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(result) == level {
		result = append(result, []int{})
	}
	result[level] = append(result[level], root.Val)
	levelOrderByDFS(root.Left, level+1)
	levelOrderByDFS(root.Right, level+1)
}

func levelOrderByBFS(root *TreeNode) [][]int {
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
