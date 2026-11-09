type simpul struct {
	panjang  int
	pref     int
	suf      int
	terbaik  int
	kiri     byte
	kanan    byte
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max2(max2(a, b), c)
}

func gabung(a, b simpul) simpul {
	if a.panjang == 0 {
		return b
	}
	if b.panjang == 0 {
		return a
	}

	c := simpul{
		panjang: a.panjang + b.panjang,
		kiri:    a.kiri,
		kanan:   b.kanan,
	}

	if a.kanan == b.kiri {
		c.terbaik = max3(a.terbaik, b.terbaik, a.suf+b.pref)
		c.pref = a.pref
		if a.pref == a.panjang {
			c.pref = a.panjang + b.pref
		}
		c.suf = b.suf
		if b.suf == b.panjang {
			c.suf = b.panjang + a.suf
		}
	} else {
		c.terbaik = max2(a.terbaik, b.terbaik)
		c.pref = a.pref
		c.suf = b.suf
	}

	return c
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	y := len(queryIndices)
	hasil := make([]int, y)
	z := len(s)

	charArray := []byte(s)
	quecharArray := []byte(queryCharacters)

	ukuran := 1
	for ukuran < z {
		ukuran <<= 1
	}

	pohon := make([]simpul, 2*ukuran)

	for i := 0; i < z; i++ {
		pohon[ukuran+i] = simpul{
			panjang: 1,
			pref:    1,
			suf:     1,
			terbaik: 1,
			kiri:    charArray[i],
			kanan:   charArray[i],
		}
	}

	for i := ukuran - 1; i >= 1; i-- {
		pohon[i] = gabung(pohon[2*i], pohon[2*i+1])
	}

	for i := 0; i < y; i++ {
		k := queryIndices[i]
		pos := ukuran + k

		pohon[pos] = simpul{
			panjang: 1,
			pref:    1,
			suf:     1,
			terbaik: 1,
			kiri:    quecharArray[i],
			kanan:   quecharArray[i],
		}

		pos >>= 1
		for pos >= 1 {
			pohon[pos] = gabung(pohon[2*pos], pohon[2*pos+1])
			pos >>= 1
		}

		hasil[i] = pohon[1].terbaik
	}

	return hasil
}
