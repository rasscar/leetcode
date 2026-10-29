func minCut(s string) int {
	var n int
	n = len(s)

	var palin [][]bool
	palin = make([][]bool, n)
	var i, j int
	for i = 0; i < n; i++ {
		palin[i] = make([]bool, n)
	}

	var akhir, awal int
	for akhir = 0; akhir < n; akhir++ {
		for awal = 0; awal <= akhir; awal++ {
			if s[awal] == s[akhir] && (akhir-awal <= 2 || palin[awal+1][akhir-1]) {
				palin[awal][akhir] = true
			}
		}
	}

	var dp []int
	dp = make([]int, n)
	for i = 0; i < n; i++ {
		if palin[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i
			for j = 0; j < i; j++ {
				if palin[j+1][i] {
					var kandidat int
					kandidat = dp[j] + 1
					if kandidat < dp[i] {
						dp[i] = kandidat
					}
				}
			}
		}
	}

	return dp[n-1]
}
