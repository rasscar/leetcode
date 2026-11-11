func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				a := dp[i][j]
				b := dp[i+1][j]
				c := dp[i][j+1]
				min := a
				if b < min {
					min = b
				}
				if c < min {
					min = c
				}
				dp[i+1][j+1] = 1 + min
			}
		}
	}
	return dp[m][n]
}
