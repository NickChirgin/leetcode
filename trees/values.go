package trees

import "sort"

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    res := []int{}
    res = append(res, getValues(root1)...)
    res = append(res, getValues(root2)...)
    sort.Ints(res)
    return res
}

func getValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    arr := []int{}
    arr = append(arr, root.Val.(int))
    arr = append(arr, getValues(root.Left)...)
    arr = append(arr, getValues(root.Right)...)
    return arr
}