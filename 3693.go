func climbStairs(n int, costs []int) int {
	var x, y, z, skr, m, i int

	x = costs[n-1]
	y = 1 << 30
	z = 1 << 30

	for i = n - 1; i >= 0; i-- {
		var biaya int
		if i == 0 {
			biaya = 0
		} else {
			biaya = costs[i-1]
		}

		m = x + 1
		if y+4 < m {
			m = y + 4
		}
		if z+9 < m {
			m = z + 9
		}

		skr = biaya + m
		z = y
		y = x
		x = skr
	}

	return x
}
