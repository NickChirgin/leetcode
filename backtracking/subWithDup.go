package backtracking

import (
	"fmt"
	"sort"
)

func SubsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := solve(nums) 
	res = append(res, []int{})
	return res
}

func solve(nums []int) [][]int{
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	res := [][]int{}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		n := nums[0]
		nums = nums[1:]
		for _, val := range solve(nums) {
			val = append(val, n)
			res = append(res, val)
		}
		res = append(res, []int{n})	
		fmt.Println(n, i)
	}
	return res
}

