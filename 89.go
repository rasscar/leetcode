func grayCode(n int) []int {
    var r[]int
    var i int
    for i = 0; i< (1<<n); i++{
        r = append(r, i^ (i>>1))
    }
    return r
}
