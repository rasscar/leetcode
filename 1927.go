func sumGame(num string) bool {
    n := len(num)
    half := n / 2
    diff := 0
    cnt := 0

    for i := 0; i < n; i++ {
        c := num[i]
        if c == '?' {
            if i < half {
                cnt++
            } else {
                cnt--
            }
        } else {
            v := int(c - '0')
            if i < half {
                diff += v
            } else {
                diff -= v
            }
        }
    }

    if cnt%2 != 0 {
        return true
    }

    return diff+cnt/2*9 != 0
}
