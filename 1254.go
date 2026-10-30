func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(i int, j int) bool
	dfs = func(i int, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if grid[i][j] == 1 {
			return true
		}
		
		grid[i][j] = 1
		
		x := dfs(i+1, j)
		y := dfs(i-1, j)
		z := dfs(i, j+1)
		w := dfs(i, j-1)
		
		return x && y && z && w
	}

	jum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				if dfs(i, j) {
					jum++
				}
			}
		}
	}
	
	return jum
}
