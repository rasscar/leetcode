func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    
    ans := temp(x, n)
    if n < 0 {
        return 1.0 / ans
    }
    return ans
}

func temp(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    
    half := temp(x, n/2)
    if n % 2 == 0 {
        return half * half
    }
    return half * half * x
}
