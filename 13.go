func romanToInt(s string) int {
    var total int
    var roman = map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i := 0; i < len(s); i++ {
        if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
            total = total - roman[s[i]]
        } else {
            total = total + roman[s[i]]
        }
    }

    return total
}
