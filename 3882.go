func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][][]bool, m)

    for i := range dp {
        dp[i] = make([][]bool, n)
        for j := range dp[i] {
            dp[i][j] = make([]bool, 1024)
        }
    }

    dp[m-1][n-1][grid[m-1][n-1]] = true

    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i == m-1 && j == n-1 {
                continue
            }

            if i+1 < m {
                for x := 0; x < 1024; x++ {
                    if dp[i+1][j][x] {
                        dp[i][j][x^grid[i][j]] = true
                    }
                }
            }

            if j+1 < n {
                for x := 0; x < 1024; x++ {
                    if dp[i][j+1][x] {
                        dp[i][j][x^grid[i][j]] = true
                    }
                }
            }
        }
    }

    for x := 0; x < 1024; x++ {
        if dp[0][0][x] {
            return x
        }
    }

    return 0
}
