package backtracking

func subsets(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
			return res
	}
	for i, val := range nums {
			res = append(res, []int{val})
			p := subsets(nums[i+1:])
			for _, v := range p {
					if len(v) != 0 {
							res = append(res, append(v, val))
					}
			}
	}
	res = append([][]int{{}}, res...)
	return res

}