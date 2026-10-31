func removeInvalidParentheses(s string) []string {
	hasil := []string{}
	var remove func(s string, scan int, del int, o byte, c byte)
	remove = func(s string, scan int, del int, o byte, c byte) {
		b := 0
		for i := scan; i < len(s); i++ {
			if s[i] == o {
				b++
			} else if s[i] == c {
				b--
			}
			if b >= 0 {
				continue
			}
			for j := del; j <= i; j++ {
				if s[j] == c && (j == del || s[j-1] != c) {
					next := s[:j] + s[j+1:]
					remove(next, i, j, o, c)
				}
			}
			return
		}
		bytesBalik := []byte(s)
		kiri := 0
		kanan := len(bytesBalik) - 1
		for kiri < kanan {
			bytesBalik[kiri], bytesBalik[kanan] = bytesBalik[kanan], bytesBalik[kiri]
			kiri++
			kanan--
		}
		balik := string(bytesBalik)
		if o == '(' {
			remove(balik, 0, 0, ')', '(')
		} else {
			hasil = append(hasil, balik)
		}
	}
	remove(s, 0, 0, '(', ')')
	return hasil
}
