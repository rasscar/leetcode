func sumCounts(nums []int) int {
	result := 0
	for i := range len(nums) {
		count := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			count[nums[j]] = true
			distinct := len(count)
			result += distinct * distinct
		}
	}
	return result
}
