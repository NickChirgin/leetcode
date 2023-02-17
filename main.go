package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/rpn"
)
func main() {
	nums := []string{"2","1","+","3","*"}
	fmt.Println(rpn.EvalRPN(nums))
}