package binarysearch

func isPerfectSquare(num int) bool {
	low := 0
	high := num
	for low <= high {
		mid := high - (high - low) / 2
		if mid * mid == num {
			return true
		} else if mid * mid > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}