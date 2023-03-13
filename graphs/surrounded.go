package graphs

func solve(board [][]byte) {
	o, t, x := []byte("O")[0], []byte("T")[0], []byte("X")[0]
	rows, cols := len(board), len(board[0])
	var dfs func(r,c int)
	dfs = func (r, c int) {
		if r < 0 || c < 0 || r == rows || c == cols || board[r][c] != o {
			return
		}
		board[r][c] = t
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	for r:=0; r< rows; r++ {
		for c:=0; c< cols; c++ {
			if board[r][c] == o && ((r == 0 || r == rows-1) || (c == 0 || c == cols -1)) {
				dfs(r,c)
			}
		}
	}
	for r:=0; r< rows; r++ {
		for c:=0; c< cols; c++ { 
			if board[r][c] == o {
				board[r][c] = x
			}
		}
	}
	for r:=0; r< rows; r++ {
		for c:=0; c< cols; c++ { 
			if board[r][c] == t {
				board[r][c] = o
			}
		}
	}
	dfs(0, 0)
}