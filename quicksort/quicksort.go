package quicksort

func Quick(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	del := nums[0]
	less := []int{}
	greater := []int{}
	for _ ,val := range nums[1:] {
		if val < del {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}
	less = Quick(less)
	greater = Quick(greater)
	less = append(less, del)
	return append(less ,greater...)
}