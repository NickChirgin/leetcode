package trees

import "sort"

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
			return 0
	}
	q := []*TreeNode{root}
	arr := []int{}
	for len(q) > 0 {
			qlen := len(q)
			for i:=0; i<qlen; i++{
					node := q[0]
					q = q[1:]
					arr = append(arr, node.Val.(int))
					if node.Left != nil {
							q = append(q, node.Left)
					}
					if node.Right != nil {
							q = append(q, node.Right)
					}
			}
	}
	sort.Ints(arr)
	return arr[k-1]
}