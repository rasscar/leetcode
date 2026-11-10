func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    min := len(strs[0])
    for i := 0; i < len(strs); i++ {
        if len(strs[i]) < min {
            min = len(strs[i])
        }
    }

    s := []byte{}

    for i := 0; i < min; i++ {
        con := true

        for j := 0; j < len(strs)-1; j++ {
            if strs[j][i] != strs[j+1][i] {
                con = false
                break
            }
        }

        if !con {
            break
        }

        s = append(s, strs[0][i])
    }

    return string(s)
}
