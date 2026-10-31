func hasValidPath(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	if (m+n-1)%2 != 0 {
		return false
	}

	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, m+n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, balance int) bool
	dfs = func(i, j, balance int) bool {
		if i >= m || j >= n {
			return false
		}

		if grid[i][j] == '(' {
			balance++
		} else {
			balance--
		}

		if balance < 0 {
			return false
		}

		if dp[i][j][balance] != -1 {
			return dp[i][j][balance] == 1
		}

		if i == m-1 && j == n-1 {
			if balance == 0 {
				dp[i][j][balance] = 1
			} else {
				dp[i][j][balance] = 0
			}
			return dp[i][j][balance] == 1
		}

		down := dfs(i+1, j, balance)
		right := dfs(i, j+1, balance)

		if down || right {
			dp[i][j][balance] = 1
		} else {
			dp[i][j][balance] = 0
		}
		return dp[i][j][balance] == 1
	}

	return dfs(0, 0, 0)
}
