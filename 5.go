func longestPalindrome(s string) string {
	var n int
	n = len(s)

	if n == 0 {
		return ""
	}

	var awal, akhir, i, mks int
	awal = 0
	akhir = 0

	var expand func(l int, r int) int
	expand = func(l int, r int) int {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		return r - l - 1
	}

	for i = 0; i < n; i++ {
		var ganjil, genap int
		ganjil = expand(i, i)
		genap = expand(i, i+1)

		mks = ganjil
		if genap > mks {
			mks = genap
		}

		if mks > akhir-awal {
			awal = i - (mks-1)/2
			akhir = i + mks/2
		}
	}

	return s[awal : akhir+1]
}
