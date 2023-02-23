package binarysearch
func findMin(nums []int) int {
	low := 0 
	high := len(nums) - 1
	min := nums[0]
	for low <= high {
			mid := (low+high)/2
			
			if nums[mid] >= nums[0] {
					low = mid + 1
			} else {
					if nums[mid] < min {
							min = nums[mid]
					}
					high = mid -1
			}
	}
	return min
}