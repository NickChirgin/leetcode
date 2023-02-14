package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/topkfrequent"
)
func main() {
	test := []int{2, 7, 11 ,15, 17, 18, 19, 20, 2, 7, 7}

	fmt.Println(topkfrequent.TopKFrequent(test, 2))
}