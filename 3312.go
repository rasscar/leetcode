func gcdValues(nums []int, queries []int64) []int {
    maxVal := 0
    for _, v := range nums {
        if v > maxVal {
            maxVal = v
        }
    }
    freq := make([]int, maxVal+1)
    for _, v := range nums {
        freq[v]++
    }
    multiples := make([]int64, maxVal+1)
    for g := 1; g <= maxVal; g++ {
        var cnt int64 = 0
        for k := g; k <= maxVal; k += g {
            cnt += int64(freq[k])
        }
        multiples[g] = cnt
    }
    exact := make([]int64, maxVal+1)
    for g := maxVal; g >= 1; g-- {
        c := multiples[g]
        exact[g] = c * (c - 1) / 2
        for k := 2 * g; k <= maxVal; k += g {
            exact[g] -= exact[k]
        }
    }
    prefix := make([]int64, maxVal+1)
    var sum int64 = 0
    for g := 1; g <= maxVal; g++ {
        sum += exact[g]
        prefix[g] = sum
    }
    ans := make([]int, len(queries))
    for i, q := range queries {
        lo, hi := 1, maxVal
        for lo < hi {
            mid := (lo + hi) / 2
            if prefix[mid] > q {
                hi = mid
            } else {
                lo = mid + 1
            }
        }
        ans[i] = lo
    }
    return ans
} 
