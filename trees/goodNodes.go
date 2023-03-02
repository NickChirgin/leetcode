package trees

func goodNodes(root *TreeNode) int {
	return goodNodesUtil(root, root.Val.(int))
}

func goodNodesUtil(root *TreeNode, parent int) int {
	if root == nil {
			return 0
	}
	
	res := 1
	max := root.Val
	
	if parent > root.Val.(int) {
			res = 0
			max = parent
	}
	
	res += goodNodesUtil(root.Left, max.(int))
	res += goodNodesUtil(root.Right, max.(int))
	
	return res
}