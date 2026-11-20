func search(nums []int, target int) int {
	var n, st, en, m int
	n = len(nums)
	st = 0
	en = n - 1

	for st <= en {
		m = st + (en-st)/2

		if nums[m] == target {
			return m
		}

		if nums[m] >= nums[st] {
			if target >= nums[st] && target < nums[m] {
				en = m - 1
			} else {
				st = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[en] {
				st = m + 1
			} else {
				en = m - 1
			}
		}
	}

	return -1
}
