package main

import (
	"fmt"

	"github.com/nickchirgin/leetcode/backtracking"
)

func main() {
	board := [][]string{{"C","A","A"},{"A","A","A"},{"B","C","D"}}
	fmt.Println(backtracking.Exist(board, "AAB"))
}
