func getPermutation(n int, k int) string {
    var f = 1
    var a []int
    var i int
    for i = 1; i <= n; i++ {
        f *= i
        a = append(a, i)
    }
    k--
    var r []string
    var j int
    for i = n; i > 0; i-- {
        f /= i
        j = k / f
        k %= f
        r = append(r, strconv.Itoa(a[j]))
        a = append(a[:j], a[j+1:]...)
    }
    return strings.Join(r, "")
}
