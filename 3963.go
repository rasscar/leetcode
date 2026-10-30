func createGrid(m int, n int) []string {
	var res [][]byte
	res = make([][]byte, m)

	var i, j int
	for i = 0; i < m; i++ {
		res[i] = make([]byte, n)
		for j = 0; j < n; j++ {
			res[i][j] = '#'
		}
	}

	var r, c int
	r = 0
	c = 0

	res[r][c] = '.'

	for r < m-1 || c < n-1 {
		var val1, val2 int
		val1 = (r+1)*(n-1) - c*(m-1)
		val2 = r*(n-1) - (c+1)*(m-1)

		var abs1, abs2 int
		abs1 = val1
		if abs1 < 0 {
			abs1 = -abs1
		}

		abs2 = val2
		if abs2 < 0 {
			abs2 = -abs2
		}

		if r < m-1 && (c == n-1 || abs1 <= abs2) {
			r++
		} else {
			c++
		}

		res[r][c] = '.'
	}

	var hasil []string
	hasil = make([]string, m)
	for i = 0; i < m; i++ {
		hasil[i] = string(res[i])
	}

	return hasil
}
