package twosum

func TwoSum(nums []int, target int) []int {
	var result []int
  for i, value := range nums {
		for j:= i+1; j < len(nums); j++ {
			if value + nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}