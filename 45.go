func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	nextHop := 0
	batasSekarang := 0
	hop := 0

	for i := 0; i < len(nums)-1; i++ {
		maksJangkau := i + nums[i]
		if maksJangkau > nextHop {
			nextHop = maksJangkau
		}
		if i == batasSekarang {
			hop++
			batasSekarang = nextHop
		}
	}
	return hop
}
