func minimumDeletions(nums []int) int {
    n := len(nums)
    minIdx, maxIdx := 0, 0

    for i, v := range nums {
        if v < nums[minIdx] {
            minIdx = i
        }
        if v > nums[maxIdx] {
            maxIdx = i
        }
    }

    a, b := minIdx, maxIdx
    if a > b {
        a, b = b, a
    }

    ans := b + 1

    if n-a < ans {
        ans = n - a
    }

    if a+1+n-b < ans {
        ans = a + 1 + n - b
    }

    return ans
}
