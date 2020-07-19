package Easy

func reserveTree(root *TreeNode, sum, count int) bool {
	count += root.Val
	if root.Left==nil&&root.Right==nil{
		if sum==count {
			return true
		}else {
			return false
		}
	}
	if root.Left!=nil&&root.Right!=nil{
		return reserveTree(root.Left,sum,count)||reserveTree(root.Right,sum,count)
	}
	if root.Left==nil&&root.Right!=nil{
		return reserveTree(root.Right,sum,count)
	}
	if root.Left!=nil&&root.Right==nil{
		return reserveTree(root.Left,sum,count)
	}
	return false
}

func PathSum(root *TreeNode, sum int) bool {
	if root==nil{
		return false
	}
	var count int
	return reserveTree(root, sum, count)
}