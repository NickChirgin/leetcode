package backtracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
			return nil
	}
	res := [][]int{}
	curr := []int{}
	sort.Ints(candidates)
	var backtracking func(curr []int, sum, idx int)
	backtracking = func(curr []int, sum, idx int) {
			if sum == target {
					res = append(res, append([]int{}, curr...))
			}
			if sum > target {
					return
			}
			for i:=idx; i<len(candidates); i++ {
					if i>idx && candidates[i]==candidates[i-1] {
							continue
					}
					curr = append(curr, candidates[i])
					backtracking(curr, sum+candidates[i], i+1)
					curr = curr[:len(curr)-1]
			}
	}
	backtracking(curr, 0, 0)
	return res
}