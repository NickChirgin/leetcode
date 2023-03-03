package backtracking

func sum(candidates []int, target int) [][]int {
	res := [][]int{}
	curr := []int{}
	var dfs func(i int, curr []int, total int)
	dfs = func (i int, curr []int, total int) {
		if total == target {
			res = append(res, append([]int{}, curr...))
			return
		}	
		if i >= len(candidates) || total > target {
			return
		}
		curr = append(curr, candidates[i])
		dfs(i, curr, total + candidates[i])
		curr = curr[:len(curr)-1]
		dfs(i+1, curr, total)
	}
	dfs(0, curr, 0)
	return res
}
