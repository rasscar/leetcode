func shortestBeautifulSubstring(s string, k int) string {
    n := len(s)
    minLen := n + 1

    left := 0
    count := 0
    for right := 0; right < n; right++ {
        if s[right] == '1' {
            count++
        }
        for count >= k {
            if count == k {
                if curLen := right - left + 1; curLen < minLen {
                    minLen = curLen
                }
            }
            if s[left] == '1' {
                count--
            }
            left++
        }
    }

    if minLen == n+1 {
        return ""
    }

    cnt := 0
    for i := 0; i < minLen; i++ {
        if s[i] == '1' {
            cnt++
        }
    }

    ans := ""
    if cnt == k {
        ans = s[:minLen]
    }

    for i := 1; i <= n-minLen; i++ {
        if s[i-1] == '1' {
            cnt--
        }
        if s[i+minLen-1] == '1' {
            cnt++
        }
        if cnt == k {
            sub := s[i : i+minLen]
            if ans == "" || sub < ans {
                ans = sub
            }
        }
    }

    return ans
}
