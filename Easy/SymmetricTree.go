package Easy

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p==nil||q==nil{
		return p==q
	}

	if p.Val!=q.Val{
		return false
	}

	return isSymmetricTree(p.Left,q.Right) && isSymmetricTree(q.Left,p.Right)
}

func SymmetricTree(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}