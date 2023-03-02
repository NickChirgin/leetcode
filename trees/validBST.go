package trees

func isValidBST(root *TreeNode) bool {
    return valid(root, nil, nil)
}

func valid(root, min, max *TreeNode) bool {
    if root == nil {
        return true
    }
    if min != nil && root.Val.(int) <= min.Val.(int) {
        return false
    }
    if max != nil && root.Val.(int) >= max.Val.(int) {
        return false
    }
    return valid(root.Left, min, root) && valid(root.Right, root, max)
}