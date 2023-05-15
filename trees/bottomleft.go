package trees

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	visited := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		visited = append(visited, node.Val.(int))
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return visited[len(visited)-1]
}