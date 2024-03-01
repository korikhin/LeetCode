package main

func IsValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	grid := make([]int, 9)

	for i, row := range board {
		for j, cell := range row {
			if cell < '1' || cell > '9' {
				continue
			}

			n := 3*(i/3) + (j / 3)
			bit := 1 << (cell - '1')

			if bit&(rows[i]|cols[j]|grid[n]) != 0 {
				return false
			}

			rows[i] |= bit
			cols[j] |= bit
			grid[n] |= bit
		}
	}

	return true
}
