func countCommas(n int64) int64 {
    ans := int64(0)
    for p := int64(1000); p <= n; {
        ans += n - p + 1
        if p > n/1000 {
            break
        }
        p *= 1000
    }

    return ans
}
