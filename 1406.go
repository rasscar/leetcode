func stoneGameIII(stoneValue []int) string {
    n := len(stoneValue)
    dp := make([]int, n+1)

    for i := n - 1; i >= 0; i-- {
        take := 0
        dp[i] = -1 << 31
        for k := 1; k <= 3 && i+k <= n; k++ {
            take += stoneValue[i+k-1]
            advantage := take - dp[i+k]
            if advantage > dp[i] {
                dp[i] = advantage
            }
        }
    }

    if dp[0] > 0 {
        return "Alice"
    } else if dp[0] < 0 {
        return "Bob"
    }
    return "Tie"
}
