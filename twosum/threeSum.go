package twosum

import (
	"sort"
	"strconv"
)
func ThreeSum(nums []int) [][]int {
    resMap := map[string]bool{}
    result := [][]int{}
    sort.Ints(nums)
    for i:=0; i <= len(nums) -3; i++ {
        firstPointer := i + 1
        secondPointer := len(nums) - 1
        for firstPointer < secondPointer {
            sum := nums[i] + nums[firstPointer] + nums[secondPointer]
            if  sum == 0 {
                key := strconv.Itoa(nums[i])+strconv.Itoa(nums[firstPointer])+strconv.Itoa(nums[secondPointer])
                if resMap[key] != true {
                    result = append(result, []int{nums[i], nums[firstPointer], nums[secondPointer]})
                    resMap[key] = true
                }     
            }
            if sum > 0 {
                secondPointer--
            } else {
                firstPointer++
            }

        }
    }
    return result
}