package backtracking

import "fmt"

func Exist(board [][]string, word string) bool {
	for i := range board {
			for j := range board[0] {
					if board[i][j] == string(word[0]) {
							fmt.Println(i, j, board[i][j])
							if check(board, word[1:], i, j) {
								return true
							}
					}
			}
	}
	return false
}

func check(board[][]string, word string, i,j int) bool {
	if len(word) == 0 {
			return true
	}
	if i + 1 < len(board) {
			if board[i+1][j] == word[:1] {
				board[i+1][j] = "*"
				check(board, word[1:], i+1, j)
				board[i+1][j] = word[:1]
			}
	}
	if j +1 < len(board[0]) {
			if board[i][j+1] == word[:1] {
				board[i][j+1] = "*"
				check(board, word[1:], i, j+1)
				board[i][j+1] = word[:1]
			}
			
	}
	if i > 0 {
			if board[i-1][j] == word[:1] {
				board[i-1][j] = "*"
				check(board, word[1:], i-1, j)
				board[i-1][j] = word[:1]
			}

	}
	if j > 0 {
			if board[i][j-1] == word[:1] {
				board[i][j-1] = "*"
				check(board, word[1:], i, j-1)
				board[i][j-1] = word[:1]
			}
	}
	return false
}
