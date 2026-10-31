func uniquePaths(grid [][]int) int {
	var mod int
	mod = 1000000007

	var down []int
	down = make([]int, len(grid[0]))
	down[0] = 1

	var i, j int
	for i = 0; i < len(grid); i++ {
		var right int
		right = 0
		for j = 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				var temp int
				temp = right
				right = down[j]
				down[j] = temp
			} else {
				var jumlah int
				jumlah = right + down[j]
				if jumlah >= mod {
					jumlah -= mod
				}
				right = jumlah
				down[j] = jumlah
			}
		}
	}
	
	return down[len(down)-1]
}
