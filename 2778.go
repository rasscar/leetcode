func sumOfSquares(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if n%(i+1) == 0 {
			ans += nums[i] * nums[i]
		}
	}
	return ans
}
