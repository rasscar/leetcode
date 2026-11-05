func lexGreaterPermutation(s string, target string) string {
    n := len(s)
    freqS := make([]int, 26)
    for _, ch := range s {
        freqS[ch-'a']++
    }
    need := make([]int, 26)
    for i := 0; i < n-1; i++ {
        need[target[i]-'a']++
    }
    for i := n - 1; i >= 0; i-- {
        validPrefix := true
        for j := 0; j < 26; j++ {
            if need[j] > freqS[j] {
                validPrefix = false
                break
            }
        }
        if validPrefix {
            targetChar := target[i] - 'a'
            for c := targetChar + 1; c < 26; c++ {
                if freqS[c] > need[c] {
                    res := make([]byte, 0, n)
                    for k := 0; k < i; k++ {
                        res = append(res, target[k])
                    }
                    res = append(res, byte('a'+c))
                    need[c]++
                    for j := 0; j < 26; j++ {
                        for k := 0; k < freqS[j]-need[j]; k++ {
                            res = append(res, byte('a'+j))
                        }
                    }
                    return string(res)
                }
            }
        }
        if i > 0 {
            need[target[i-1]-'a']--
        }
    }
    return ""
}
