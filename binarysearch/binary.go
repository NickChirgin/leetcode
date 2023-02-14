package binarysearch

func Binary(nums []int, target int) int {
	high := int(len(nums)-1)
	low := 0
	for low<=high{
		mid := int((low + high)/2)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}