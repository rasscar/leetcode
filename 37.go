func solveSudoku(b [][]byte) {
	var r, c, bx [9][10]int
	var solved bool
	solved = false

	var place func(d int, r int, c int)
	place = func(d int, r int, c int) {
		var idx int
		idx = (r/3)*3 + c/3
		r[r][d]++
		c[c][d]++
		bx[idx][d]++
		b[r][c] = byte(d + '0')
	}

	var remove func(d int, r int, c int)
	remove = func(d int, r int, c int) {
		var idx int
		idx = (r/3)*3 + c/3
		r[r][d]--
		c[c][d]--
		bx[idx][d]--
		b[r][c] = '.'
	}

	var next func(row int, col int)
	next = func(row int, col int) {
		if row == 8 && col == 8 {
			solved = true
		} else if col == 8 {
			backtrack(row+1, 0)
		} else {
			backtrack(row, col+1)
		}
	}

	var backtrack func(row int, col int)
	backtrack = func(row int, col int) {
		if b[row][col] == '.' {
			for d := 1; d <= 9; d++ {
				var idx int
				idx = (row/3)*3 + col/3
				if r[row][d]+c[col][d]+bx[idx][d] == 0 {
					place(d, row, col)
					next(row, col)
					if !solved {
						remove(d, row, col)
					}
				}
			}
		} else {
			next(row, col)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] != '.' {
				var d int
				d = int(b[i][j] - '0')
				place(d, i, j)
			}
		}
	}

	backtrack(0, 0)
}
