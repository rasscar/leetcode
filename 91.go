func cekValid(kode int, tipe int) bool {
	if tipe == 1 {
		return kode >= 1 && kode <= 26
	}
	return kode >= 10 && kode <= 26
}

func numDecodings(s string) int {
	var n int
	n = len(s)

	if n == 0 {
		return 0
	}

	var awal, skr, hasil, i int
	awal = 1
	skr = 0

	if cekValid(int(s[0]-'0'), 1) {
		skr = 1
	}

	for i = 2; i <= n; i++ {
		hasil = 0
		if cekValid(int(s[i-1]-'0'), 1) {
			hasil += skr
		}
		if cekValid(int(s[i-2]-'0')*10+int(s[i-1]-'0'), 2) {
			hasil += awal
		}
		awal = skr
		skr = hasil
	}

	return skr
}
