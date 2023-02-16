package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/maxArea"
	"github.com/nickchirgin/leetcode/validParenthesis"
)
func main() {
	nums := []int{1,3,2,5,25,24,5}
	fmt.Println(maxArea.MaxArea(nums))
	fmt.Println(validParenthesis.ValidParenthesis("()[]{}"))
}