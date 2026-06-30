func isValidSudoku(board [][]byte) bool {
	row:=len(board)
	if row == 0 {
		return true
	}
	cols:=[9][9]bool{}
	rows:=[9][9]bool{}
	boxes:=[9][9]bool{}
	for r:=range board{
		for c:=range board[r]{
			if board[r][c] == '.'{
				continue
			}
			bi:=(r/3)+(c/3)*3
			n:=board[r][c]-'0'-1
			if rows[r][n] || cols[c][n] || boxes[bi][n] {
				return false
			}
			rows[r][n] = true 
			cols[c][n] = true 
			boxes[bi][n] = true
		}
	}
	return true
}
