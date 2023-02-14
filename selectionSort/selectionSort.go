package selectionSort

func findSmallest(nums []int) int{
	smallest := nums[0]
	smallestIndex := 0
	for i, num := range nums {
		if num < smallest {
			smallest = num
			smallestIndex = i
		}
	}
	return smallestIndex
}

func SelectionSort(nums []int) []int {
	result := make([]int, len(nums))
	for i := range nums {
		index := findSmallest(nums)
		result[i] = nums[index]
		nums = append(nums[:index], nums[index+1:]...)
	}
return result
}