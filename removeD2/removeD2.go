package removeD2

func removeDuplicates(nums []int) int {
	count := 0
	l, r := 0, 1
	if len(nums) < 3 {
		return len(nums)
	}
	for r < len(nums) {
		if nums[l] == nums[r] {
			count++
		} else {
			count = 0
		}
		if count >= 2 {
			nums = append(nums[:r], nums[r+1:]...)
			continue
		}
		l++
		r++
	}
	return len(nums)
}