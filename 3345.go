func digitProduct(x int) int {
    if x == 0 {
        return 0
    }
    p := 1
    for x > 0 {
        p *= x % 10
        x /= 10
    }
    return p
}

func smallestNumber(n int, t int) int {
    for x := n; ; x++ {
        if digitProduct(x)%t == 0 {
            return x
        }
    }
}
