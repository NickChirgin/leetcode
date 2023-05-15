package backtracking

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	var backtracking func(a int, arr []int)
	backtracking = func(a int, arr []int) {
		if len(arr) == 4  && arrSum(arr) == target {
			res = append(res, arr)
		}

	}
}

func arrSum(nums []int) int {
	result := 0
	for _, val := range nums {
		result += val
	}
	return result
}