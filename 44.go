func isMatch(s string, p string) bool {
	dp := make([]bool, len(p)+1)
	dp[0] = true
	for j := 0; j < len(p) && p[j] == '*'; j++ {
		dp[j+1] = true
	}

	for i := 1; i <= len(s); i++ {
		dp[0] = false
		barisLalu := i == 1
		for j := 1; j <= len(p); j++ {
			simpanLama := dp[j]
			if p[j-1] == '*' {
				dp[j] = simpanLama || dp[j-1]
			} else {
				dp[j] = (s[i-1] == p[j-1] || p[j-1] == '?') && barisLalu
			}
			barisLalu = simpanLama
		}
	}

	return dp[len(p)]
}
