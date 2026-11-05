func distinctSubseqII(s string) int {
    const mod = 1000000007
    runes := []rune(s)
    dp := []int{1}
    last := make(map[rune]int)
    for i, ch := range runes {
        newVal := (dp[len(dp)-1] * 2) % mod
        if idx, ok := last[ch]; ok {
            newVal = (newVal - dp[idx] + mod) % mod
        }
        dp = append(dp, newVal)
        last[ch] = i
    }
    return (dp[len(dp)-1] - 1 + mod) % mod
}
