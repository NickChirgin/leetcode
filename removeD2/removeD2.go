package removeD2

func removeDuplicates(nums []int) int {
	set := map[int]int{}
	i:= 0 
	for i < len(nums) {
			set[nums[i]]++
			if set[nums[i]] > 2 {
					nums = append(nums[:i], nums[i+1:]...)
					continue
			}
			i++
	}
	return len(nums)
}