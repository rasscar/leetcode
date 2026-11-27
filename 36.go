func isValidSudoku(b [][]byte) bool {
	var r, c, bx [][]bool
	r = make([][]bool, 9)
	c = make([][]bool, 9)
	bx = make([][]bool, 9)

	var i, j, idx int
	var v byte

	for i = 0; i < 9; i++ {
		r[i] = make([]bool, 9)
		c[i] = make([]bool, 9)
		bx[i] = make([]bool, 9)
	}

	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if b[i][j] != '.' {
				v = b[i][j] - '1'
				idx = (i/3)*3 + (j/3)

				if r[i][v] || c[j][v] || bx[idx][v] {
					return false
				}
				
				r[i][v] = true
				c[j][v] = true
				bx[idx][v] = true
			}
		}
	}

	return true
}
