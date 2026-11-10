func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount(nums []int) int {
	var mod int
	mod = 1000000007

	var n, i, j, k int
	n = len(nums)

	var dp [][][]int
	dp = make([][][]int, n)
	for i = 0; i < n; i++ {
		dp[i] = make([][]int, 201)
		for j = 0; j < 201; j++ {
			dp[i][j] = make([]int, 201)
			for k = 0; k < 201; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(idx int, g1 int, g2 int) int
	dfs = func(idx int, g1 int, g2 int) int {
		if idx == n {
			if g1 != 0 && g1 == g2 {
				return 1
			}
			return 0
		}

		if dp[idx][g1][g2] != -1 {
			return dp[idx][g1][g2]
		}

		var ans int
		ans = dfs(idx+1, g1, g2)

		var ng1 int
		if g1 == 0 {
			ng1 = nums[idx]
		} else {
			ng1 = gcd(g1, nums[idx])
		}
		ans = (ans + dfs(idx+1, ng1, g2)) % mod

		var ng2 int
		if g2 == 0 {
			ng2 = nums[idx]
		} else {
			ng2 = gcd(g2, nums[idx])
		}
		ans = (ans + dfs(idx+1, g1, ng2)) % mod

		dp[idx][g1][g2] = ans
		return ans
	}

	return dfs(0, 0, 0)
}
