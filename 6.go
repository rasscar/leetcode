func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	baris := make([][]byte, numRows)
	idx := 0
	arah := 1
	for i := 0; i < len(s); i++ {
		c := s[i]
		baris[idx] = append(baris[idx], c)
		if idx == 0 {
			arah = 1
		} else if idx == numRows-1 {
			arah = -1
		}
		idx += arah
	}
	hasil := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(baris[i]); j++ {
			hasil = append(hasil, baris[i][j])
		}
	}
	return string(hasil)
}
