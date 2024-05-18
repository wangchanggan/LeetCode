package binary_tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	var child, parent *TreeNode = root, nil
	for child != nil && child.Val != key {
		parent = child
		if child.Val > key {
			child = child.Left
		} else {
			child = child.Right
		}
	}
	// not found key treeNode
	if child == nil {
		return root
	}
	if child.Left == nil && child.Right == nil {
		child = nil
	} else if child.Right == nil {
		child = child.Left
	} else if child.Left == nil {
		child = child.Right
	} else {
		r, rParent := child.Right, child
		for r.Left != nil {
			rParent = r
			r = r.Left
		}
		if rParent.Val == child.Val {
			rParent.Right = r.Right
		} else {
			rParent.Left = r.Right
		}
		r.Right = child.Right
		r.Left = child.Left
		child = r
	}
	if parent == nil {
		return child
	}
	if parent.Left != nil && parent.Left.Val == key {
		parent.Left = child
	} else {
		parent.Right = child
	}
	return root
}
