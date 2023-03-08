package trees

func hasPathSum(root *TreeNode, targetSum int) bool {
    var dfs func(node *TreeNode, sum int) bool
    dfs = func(node *TreeNode, sum int) bool {
        if node == nil {
            return false
        }
        sum += node.Val.(int)
        if node.Left == nil && node.Right == nil {
            return sum==targetSum
        }
        return dfs(node.Left, sum) || dfs(node.Right, sum)
    }
    return dfs(root, 0)
}