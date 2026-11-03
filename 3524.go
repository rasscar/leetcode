func resultArray(nums []int, k int) []int64 {
	n := int64(len(nums))
	if k == 1 {
		return []int64{n * (n + 1) / 2}
	}

	res := make([]int64, k)
	dp := make([]int64, k)

	for _, v := range nums {
		m := v % k
		dp2 := make([]int64, k)
		
		dp2[0] = dp[0]
		for i := 1; i < k; i++ {
			idx := (i * m) % k
			val := dp[i]
			if i == 1 {
				val++
			}
			dp2[idx] += val
		}
		
		for idx := 0; idx < k; idx++ {
			res[idx] += dp2[idx]
		}
		
		dp = dp2
	}
	return res
}
