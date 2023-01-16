package array

func IsValidSudoku(board [][]byte) bool {

	for i := 0; i < len(board); i++ {
		res := []byte{}
		for j := 0; j < len(board[i]); j++ {
			res = append(res, board[j][i])
		}
		if !checkRow(board[i]) || !checkRow(res) {
			return false
		}
	}

	for i := 0; i < len(board); i += 3 {

		for j := 0; j < len(board); j += 3 {
			if !checkGrid(board, i, j) {
				return false
			}
		}
	}
	return true
}

func checkRow(row []byte) bool {
	check := make(map[byte]bool)

	for i := 0; i < 9; i++ {
		if row[i] != 46 {
			if check[row[i]] {
				return false
			} else {
				check[row[i]] = true
			}
		}

	}
	return true
}

func checkGrid(nums [][]byte, startRow int, startCol int) bool {
	check := make(map[byte]bool)

	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if nums[i][j] != 46 {
				if check[nums[i][j]] {
					return false
				} else {
					check[nums[i][j]] = true
				}

			}
		}
	}

	return true
}
