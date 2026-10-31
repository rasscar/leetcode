func maxAlternatingSum(nums []int) int64 {
    sort.Ints(nums)
    i, j := 0, len(nums)-1
    count := len(nums)/2
    if len(nums)%2 == 1 {
        count++
    }
    var ans int64 = 0
    for count > 0 {
        ai, aj := nums[i], nums[j]
        if abs(ai) >= abs(aj) {
            ans += int64(ai) * int64(ai)
            i++
        } else {
            ans += int64(aj) * int64(aj)
            j--
        }
        count--
    }
    for k := i; k <= j; k++ {
        ans -= int64(nums[k]) * int64(nums[k])
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
