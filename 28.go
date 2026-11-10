func strStr(haystack string, needle string) int {
    n, m := len(haystack), len(needle)
    if m == 0 {
        return 0
    }
    for i := 0; i <= n-m; i++ {
        match := true
        for j := 0; j < m; j++ {
            if haystack[i+j] != needle[j] {
                match = false
                break
            }
        }
        if match {
            return i
        }
    }
    return -1
}
