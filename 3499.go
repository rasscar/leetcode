func maxActiveSectionsAfterTrade(s string) int {
    ones := 0
    l := []int{}
    t := []byte{}
    for i := 0; i < len(s); {
        j := i
        for j < len(s) && s[j] == s[i] { j++ }
        l = append(l, j-i)
        t = append(t, s[i])
        if s[i] == '1' { ones += j - i }
        i = j
    }
    maxG := 0
    for i := 1; i+1 < len(l); i++ {
        if t[i] == '1' {
            g := l[i-1] + l[i+1]
            if g > maxG { maxG = g }
        }
    }
    return ones + maxG
}
