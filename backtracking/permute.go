package backtracking

func Permute(nums []int) [][]int {
	curr := []int{}
	res := [][]int{}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for _, v := range nums {
			val := nums[0]
			nums = nums[1:]
			curr = append(curr, v)
			perm := Permute(nums)
			for _, t := range perm {
				t = append(t, val)
				res = append(res, t)
			}
			nums = append(nums, val)
		}
	return res
}




