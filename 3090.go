func maximumLengthSubstring(s string) int {
	frek := [26]int{}
	kiri := 0
	maks := 0

	for kanan := 0; kanan < len(s); kanan++ {
		idx := s[kanan] - 'a'
		frek[idx]++

		for frek[idx] > 2 {
			idxKiri := s[kiri] - 'a'
			frek[idxKiri]--
			kiri++
		}

		panjang := kanan - kiri + 1
		if panjang > maks {
			maks = panjang
		}
	}
	return maks
}
