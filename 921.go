func minAddToMakeValid(s string) int {
    needOpen := 0
    balance := 0

    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            balance++
        } else {
            if balance > 0 {
                balance--
            } else {
                needOpen++
            }
        }
    }

    return needOpen + balance
}
