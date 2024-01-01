package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = binaryTreeInorderTraversal(root, res)
	return res
}

func binaryTreeInorderTraversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = binaryTreeInorderTraversal(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = binaryTreeInorderTraversal(root.Right, res)
	}
	return res
}
