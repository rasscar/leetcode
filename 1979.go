func findGCD(nums []int) int {
    minVal, maxVal := nums[0], nums[0]
    for _, v := range nums {
        if v < minVal {
            minVal = v
        }
        if v > maxVal {
            maxVal = v
        }
    }
    return gcd(minVal, maxVal)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}
