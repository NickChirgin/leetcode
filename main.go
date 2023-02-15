package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/twosum"
	"github.com/nickchirgin/leetcode/validPalyndrome"
)
func main() {
	nums := []int{2, 7, 11, 15, 25}
	fmt.Println(validPalyndrome.ValidPalidrome("A man, a plan, a canal: Panama"))
	fmt.Println(twosum.TwoSumPointer(nums, 26))
}