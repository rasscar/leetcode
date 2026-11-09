func predictTheWinner(nums []int) bool {
    n := len(nums)
    dp := make([]int, n)
    copy(dp, nums)
    for length := 2; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1
            left := nums[i] - dp[i+1]
            right := nums[j] - dp[i]
            if left > right {
                dp[i] = left
            } else {
                dp[i] = right
            }
        }
    }
    return dp[0] >= 0
}
