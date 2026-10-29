func partition(s string) [][]string {
	var n int
	n = len(s)
	var hasil [][]string
	hasil = make([][]string, 0)
	var jalur []string
	jalur = make([]string, 0)

	var cekPalin func(kiri int, kanan int) bool
	cekPalin = func(kiri int, kanan int) bool {
		for kiri < kanan {
			if s[kiri] == s[kanan] {
				kiri++
				kanan--
			} else {
				return false
			}
		}
		return true
	}

	var dfs func(ind int)
	dfs = func(ind int) {
		if ind == n {
			var salin []string
			salin = make([]string, len(jalur))
			var i int
			for i = 0; i < len(jalur); i++ {
				salin[i] = jalur[i]
			}
			hasil = append(hasil, salin)
			return
		}

		var i int
		for i = ind; i < n; i++ {
			if cekPalin(ind, i) {
				jalur = append(jalur, s[ind:i+1])
				dfs(i + 1)
				jalur = jalur[:len(jalur)-1]
			}
		}
	}

	dfs(0)
	return hasil
}
