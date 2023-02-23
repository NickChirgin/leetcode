package binarysearch
func search(nums []int, target int) int {
	low := 0 
	high := len(nums) - 1
	for low <= high {
			mid := (low+high)/2
			if nums[mid] == target {
					return mid
			}
		 if nums[low] <= nums[mid] {
					if target > nums[mid] || target < nums[low] {
							low = mid + 1
					} else {
							high = mid - 1
					}
			// Right sorted portion
			} else {
					if target < nums[mid] || target > nums[high] {
							high = mid - 1
					} else {
							low = mid + 1
					}
			}
	}
	return -1
}