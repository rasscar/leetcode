func maxSum(nums []int, k int) int {
    const MOD int64 = 1_000_000_007
    bits := make([]int, 31)

    for _, x := range nums {
        for b := 0; b < 31; b++ {
            if x&(1<<b) != 0 {
                bits[b]++
            }
        }
    }

    var ans int64

    for ; k > 0; k-- {
        var x int64

        for b := 0; b < 31; b++ {
            if bits[b] > 0 {
                x |= 1 << b
                bits[b]--
            }
        }

        ans = (ans + x*x) % MOD
    }

    return int(ans)
}
