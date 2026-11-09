func stoneGameII(piles []int) int {
    n := len(piles)
    jumlahAkhir := make([]int, n+1)
    for i := n - 1; i >= 0; i-- {
        jumlahAkhir[i] = jumlahAkhir[i+1] + piles[i]
    }

    memo := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        memo[i] = make([]int, n+1)
        for j := 0; j <= n; j++ {
            memo[i][j] = -1
        }
    }

    var dfs func(i, m int) int
    dfs = func(i, m int) int {
        if i == n {
            return 0
        }
        if memo[i][m] != -1 {
            return memo[i][m]
        }

        terbaik := 0
        batas := 2 * m
        if batas > n-i {
            batas = n - i
        }

        for x := 1; x <= batas; x++ {
            mBaru := m
            if x > mBaru {
                mBaru = x
            }
            nilai := jumlahAkhir[i] - dfs(i+x, mBaru)
            if nilai > terbaik {
                terbaik = nilai
            }
        }

        memo[i][m] = terbaik
        return terbaik
    }

    return dfs(0, 1)
}
