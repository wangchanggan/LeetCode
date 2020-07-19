package Easy

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right+ 1
	}
}

func reserveBinaryTree(root *TreeNode, num [][]int, count int){
	if root!=nil{
		num[count] = append(num[count], root.Val)
		count++
	}else {
		return
	}
	reserveBinaryTree(root.Left, num, count)
	reserveBinaryTree(root.Right, num, count)

}

func BinaryTreeLevelOrderTraversalII(root *TreeNode) [][]int {
	num := make([][]int, maxDepth(root))
	result := make([][]int, maxDepth(root))

	count := 0
	if root!=nil{
		num[count] = append(num[count], root.Val)
		count++
	}else {
		return num
	}
	reserveBinaryTree(root.Left, num, count)
	reserveBinaryTree(root.Right, num, count)
	l := len(num)-1
	for i:=0;i<len(num);i++{
		result[i] = num[l]
		l--
	}
	return result
}