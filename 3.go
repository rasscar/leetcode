func lengthOfLongestSubstring(s string) int {
	var n, kiri, i, mks, pjg int
	n = len(s)
	kiri = 0
	mks = 0

	var ada [256]bool

	for i = 0; i < n; i++ {
		for ada[s[i]] {
			ada[s[kiri]] = false
			kiri++
		}
		ada[s[i]] = true

		pjg = i - kiri + 1
		if pjg > mks {
			mks = pjg
		}
	}

	return mks
}
