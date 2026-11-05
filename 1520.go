func maxNumOfSubstrings(s string) []string {
    charPosMp := make(map[byte][]int)
    for i := 0; i < len(s); i++ {
        char := s[i] - 'a'
        v, ok := charPosMp[char]
        if !ok {
            v = []int{i, i}
        } else {
            v[1] = i
        }
        charPosMp[char] = v
    }

    validInterval := [][]int{}

    for _, pos := range charPosMp {
        start := pos[0]
        end := pos[1]
        isValid := true
        for i := start; i <= end; i++ {
            tmp := charPosMp[s[i]-'a']
            if tmp[0] < start {
                isValid = false
                break
            }
            if tmp[1] > end {
                end = tmp[1]
            }
        }
        if !isValid {
            continue
        }
        validInterval = append(validInterval, []int{start, end})
    }

    flag := make([]int, len(validInterval))
    for i, interval := range validInterval {
        if flag[i] == 1 {
            continue
        }
        for j := 0; j < len(validInterval); j++ {
            if j == i {
                continue
            }
            if flag[j] == 1 {
                continue
            }
            tmp := validInterval[j]
            if tmp[0] < interval[0] && tmp[1] > interval[1] {
                flag[j] = 1
            }
        }
    }

    ret := []string{}
    for i := 0; i < len(flag); i++ {
        if flag[i] == 1 {
            continue
        }
        start, end := validInterval[i][0], validInterval[i][1]
        ret = append(ret, s[start:end+1])
    }
    return ret
}
