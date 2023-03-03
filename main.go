package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/backtracking"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(backtracking.SubsetsWithDup(nums))
}
