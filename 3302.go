func validSequence(word1 string, word2 string) []int {
	kata1 := word1
	kata2 := word2
	panj1, panj2 := len(kata1), len(kata2)
	if panj2 == 0 {
		return []int{}
	}

	kemunculan := make([][]int, 26)
	for i := 0; i < panj1; i++ {
		c := kata1[i] - 'a'
		kemunculan[c] = append(kemunculan[c], i)
	}

	rakus := make([]int, panj2)
	sebelumnya := -1
	indeksGagal := panj2
	for j := 0; j < panj2; j++ {
		c := kata2[j] - 'a'
		arr := kemunculan[c]
		k := sort.SearchInts(arr, sebelumnya+1)
		if k == len(arr) {
			rakus[j] = -1
			if indeksGagal == panj2 {
				indeksGagal = j
			}
		} else {
			rakus[j] = arr[k]
			sebelumnya = arr[k]
		}
	}
	sempurna := indeksGagal == panj2

	akhiran := make([]int, panj1+1)
	s := 0
	for i := panj1 - 1; i >= 0; i-- {
		if s < panj2 && kata1[i] == kata2[panj2-1-s] {
			s++
		}
		akhiran[i] = s
	}
	akhiran[panj1] = 0

	type PohonFenwick struct {
		n   int
		bit []int
	}
	buatPohon := func(n int) *PohonFenwick {
		return &PohonFenwick{n: n, bit: make([]int, n+1)}
	}
	tambah := func(p *PohonFenwick, idx int) {
		for i := idx; i <= p.n; i += i & -i {
			p.bit[i]++
		}
	}
	jumlah := func(p *PohonFenwick, idx int) int {
		if idx <= 0 {
			return 0
		}
		if idx > p.n {
			idx = p.n
		}
		res := 0
		for i := idx; i > 0; i -= i & -i {
			res += p.bit[i]
		}
		return res
	}
	ke := func(p *PohonFenwick, k int) int {
		idx := 0
		bit := 1
		for bit<<1 <= p.n {
			bit <<= 1
		}
		for ; bit > 0; bit >>= 1 {
			nxt := idx + bit
			if nxt <= p.n && p.bit[nxt] < k {
				idx = nxt
				k -= p.bit[nxt]
			}
		}
		return idx + 1
	}

	pohon := make([]*PohonFenwick, 26)
	for i := 0; i < 26; i++ {
		pohon[i] = buatPohon(panj1)
	}
	aktif := make([]int, 26)

	urutan := make([]int, panj1)
	for i := 0; i < panj1; i++ {
		urutan[i] = i
	}
	sort.Slice(urutan, func(i, j int) bool {
		return akhiran[urutan[i]+1] > akhiran[urutan[j]+1]
	})
	penunjuk := 0

	sah := make([]bool, panj2)
	indeksTidakCocok := make([]int, panj2)

	batas := indeksGagal
	if batas >= panj2 {
		batas = panj2 - 1
	}

	for j := 0; j <= batas && j < panj2; j++ {
		butuhAkhiran := panj2 - j - 1

		for penunjuk < panj1 && akhiran[urutan[penunjuk]+1] >= butuhAkhiran {
			idx := urutan[penunjuk]
			c := kata1[idx] - 'a'
			tambah(pohon[c], idx+1)
			aktif[c]++
			penunjuk++
		}

		sebelumnyaIdx := -1
		if j > 0 {
			sebelumnyaIdx = rakus[j-1]
		}
		target := kata2[j] - 'a'

		terbaik := panj1
		for c := 0; c < 26; c++ {
			if c == int(target) {
				continue
			}
			if aktif[c] == 0 {
				continue
			}
			p := pohon[c]
			cnt := jumlah(p, sebelumnyaIdx+1)
			if cnt < aktif[c] {
				idx := ke(p, cnt+1) - 1
				if idx < terbaik {
					terbaik = idx
				}
			}
		}

		if terbaik < panj1 {
			sah[j] = true
			indeksTidakCocok[j] = terbaik
		}
	}

	nanti := make([]bool, panj2+1)
	nanti[panj2] = sempurna
	for j := panj2 - 1; j >= 0; j-- {
		nanti[j] = nanti[j+1]
		if j+1 < panj2 && sah[j+1] {
			nanti[j] = true
		}
	}

	jawaban := make([]int, 0, panj2)

	for j := 0; j < panj2; j++ {
		exactMungkin := nanti[j]
		exactNilai := -1
		if j < indeksGagal {
			exactNilai = rakus[j]
		}

		tidakCocokMungkin := sah[j]
		tidakCocokNilai := -1
		if tidakCocokMungkin {
			tidakCocokNilai = indeksTidakCocok[j]
		}

		if exactMungkin && exactNilai != -1 && (!tidakCocokMungkin || exactNilai < tidakCocokNilai) {
			jawaban = append(jawaban, exactNilai)
			continue
		}

		if tidakCocokMungkin {
			hasil := make([]int, 0, panj2)
			if j > 0 {
				hasil = append(hasil, rakus[:j]...)
			}
			hasil = append(hasil, tidakCocokNilai)

			pos := tidakCocokNilai
			for q := j + 1; q < panj2; q++ {
				c := kata2[q] - 'a'
				arr := kemunculan[c]
				k := sort.SearchInts(arr, pos+1)
				if k == len(arr) {
					return []int{}
				}
				idx := arr[k]
				hasil = append(hasil, idx)
				pos = idx
			}
			return hasil
		}

		return []int{}
	}

	return jawaban
}
