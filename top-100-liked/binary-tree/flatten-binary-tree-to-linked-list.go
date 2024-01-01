package binary_tree

func flatten(root *TreeNode) {
	r1 := root
	res := preorderTraversal(r1)
	r2 := root
	for i := 0; i < len(res); i++ {
		r2.Val = res[i]
		r2.Left = nil
		if i != len(res)-1 {
			r2.Right = new(TreeNode)
		}
		r2 = r2.Right
	}
}
