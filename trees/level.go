package trees

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
			return [][]int{}
	}
	res := [][]int{{}}
	q := []*TreeNode{root}
	count := 0
	for len(q) > 0 {
			curlen := len(q)
			for i:= 0; i < curlen; i++ {
					node := q[0]
					q = q[1:]
					if len(res) <= count {
							res = append(res, []int{node.Val.(int)})
					} else {
							res[count] = append(res[count], node.Val.(int))
					}
					if node.Left != nil {
							q = append(q, node.Left)
					}
					if node.Right != nil {
							q = append(q, node.Right)
					}
			}
			count++
	}
	return res
}