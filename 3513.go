func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	m := 0
	for t := n; t > 0; t >>= 1 {
		m++
	}
	return 1 << m
}
