func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n == 1 {
		return 0
	}

	pref := make([]int64, n+1)
	for i, v := range stoneValue {
		pref[i+1] = pref[i] + int64(v)
	}

	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}

	const negInf = int64(-1) << 60

	p := make([]int, n)
	leftMax := make([]int64, n)

	for i := 0; i < n; i++ {
		p[i] = i
		leftMax[i] = negInf
	}

	for j := 1; j < n; j++ {
		suffixMax := make([]int64, n+2)
		for idx := range suffixMax {
			suffixMax[idx] = negInf
		}
		suffixMax[j] = -pref[j]

		for i := j - 1; i >= 0; i-- {
			S := pref[j+1] - pref[i]

			for p[i] < j && 2*(pref[p[i]+1]-pref[i]) < S {
				val := dp[i][p[i]] + pref[p[i]+1]
				if val > leftMax[i] {
					leftMax[i] = val
				}
				p[i]++
			}

			best := negInf

			if p[i] > i {
				if val := leftMax[i] - pref[i]; val > best {
					best = val
				}
			}

			if p[i] < j && 2*(pref[p[i]+1]-pref[i]) == S {
				sumL := pref[p[i]+1] - pref[i]
				val := dp[i][p[i]]
				if dp[p[i]+1][j] > val {
					val = dp[p[i]+1][j]
				}
				val += sumL
				if val > best {
					best = val
				}
			}

			if p[i]+2 <= j && suffixMax[p[i]+2] != negInf {
				if val := suffixMax[p[i]+2] + pref[j+1]; val > best {
					best = val
				}
			}

			if p[i] < j && 2*(pref[p[i]+1]-pref[i]) > S {
				sumR := pref[j+1] - pref[p[i]+1]
				if val := dp[p[i]+1][j] + sumR; val > best {
					best = val
				}
			}

			dp[i][j] = best

			val := dp[i][j] - pref[i]
			if val > suffixMax[i+1] {
				suffixMax[i] = val
			} else {
				suffixMax[i] = suffixMax[i+1]
			}
		}
	}

	return int(dp[0][n-1])
}
