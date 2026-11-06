func checkDivisibility(n int) bool {
    var a,b,c,d int
    var e bool
    a = n
    d = a % 10
    b = 1
    for true{
        if a == 0 {
            break
        }
        c = c + d
        b = b * d
        a = a / 10
        d = a % 10
    }
    e = false
    if n % (b+c) == 0{
        e = true
    }
    return e
}
