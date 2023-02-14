package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/binarysearch"
	"github.com/nickchirgin/leetcode/twosum"
)
func main() {
	test := []int{2, 7, 11 ,15, 17, 18, 19, 20, 25}
	fmt.Println(twosum.TwoSum(test, 9))
	fmt.Println(binarysearch.Binary(test, 25))
}