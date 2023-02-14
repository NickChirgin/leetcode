package exceptSelf
	
func ProductExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))
    resultMap := make(map[int]int)
    for i, v := range nums {
        if _, ok := resultMap[v]; !ok {
            resultMap[v] = multiply(nums[:i]) * multiply(nums[i+1:])
        } 
        answer[i] = resultMap[v]
    }
    return answer
}
func multiply(nums []int) int {
    res := 1
    for _, value := range nums {
        res *= value
    }
    return res
}