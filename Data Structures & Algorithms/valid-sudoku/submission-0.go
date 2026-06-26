func isValidSudoku(board [][]byte) bool {
    if len(board) == 0 {
        return true
    }
    dot := byte('.')
    ln := 9
    for i := range board {
        set := map[byte]bool{}
        for j := range board[i] {
            if board[i][j] == dot {
                continue
            }
            if set[board[i][j]] {
                return false
            }
            set[board[i][j]]=true
        }
    }
    for i := range ln{
        set:= map[byte]bool{}
        for j:=range board {
            if board[j][i] == dot {
                continue
            }
            if set[board[j][i]]{
                return false
            }
            set[board[j][i]]=true
        } 
    }
    for square := range ln {
        set := map[byte]bool{}
        for i:= range 3{
            for j:= range 3 {
              row := (square/3)*3 + i
              col := (square%3)*3 + j
              if board[row][col] == dot {
                continue
              }  
              if set[board[row][col]] {
                return false
              }
              set[board[row][col]]=true
            }
        }
    }
    return true
}
