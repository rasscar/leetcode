func longestMonotonicSubarray(nums []int) int {
	var n, i, mks, naik, turun, temp int
	n = len(nums)
	mks = 1
	naik = 1
	turun = 1

	for i = 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			naik++
			turun = 1
		} else if nums[i] < nums[i-1] {
			turun++
			naik = 1
		} else {
			naik = 1
			turun = 1
		}

		temp = naik
		if turun > temp {
			temp = turun
		}
		if temp > mks {
			mks = temp
		}
	}

	return mks
}
