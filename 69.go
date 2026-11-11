func mySqrt(n int) int {
 if n < 2 {
        return n
    }
    x := n
    y := (x + 1) / 2
    for y < x {
        x = y
        y = (x + n/x) / 2
    }
    return x
}
