package binarysearch

func SearchMatrix(matrix [][]int, target int) bool {
	low := 0
	high := len(matrix)*len(matrix[0]) - 1
	for low <= high {
		mid := (low + high)/2
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}