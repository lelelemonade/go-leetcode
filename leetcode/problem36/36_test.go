package problem36

func isValidSudoku(board [][]byte) bool {
	var rows [9][9]int
	var columns [9][9]int
	var subBoxs [3][3][9]int

	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			val:=board[i][j]
			if val == '.'{
				continue
			}
			rows[i][val-'1']++
			columns[j][val-'1']++
			subBoxs[i/3][j/3][val-'1']++

			if rows[i][val-'1']>1||columns[j][val-'1']>1||subBoxs[i/3][j/3][val-'1']>1{
				return false
			}
		}
	}
	return true
}