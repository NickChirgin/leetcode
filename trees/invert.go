package trees

type TreeNode struct {
	Val any
	Left *TreeNode
	Right *TreeNode
}
func Invert(root * TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	Invert(root.Left)
	Invert(root.Right)
	return root
}