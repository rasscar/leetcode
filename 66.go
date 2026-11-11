func plusOne(digits []int) []int {
    var n, carry, i int
    n = len(digits)
    carry = 1

    for i = n - 1; i >= 0; i-- {
        if digits[i]+carry == 10 {
            digits[i] = 0
            carry = 1
        } else {
            digits[i] = digits[i] + carry
            carry = 0
            break
        }
    }

    if carry == 1 {
        var result []int
        result = append(result, 1)
        result = append(result, digits...)
        return result
    }

    return digits
}
