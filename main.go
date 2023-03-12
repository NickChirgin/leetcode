package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/heaps"
)

func main() {
	board := [][]int{{3, 2},{1, 3},{2, 6}}
	fmt.Println(heaps.GetOrder(board))
}
