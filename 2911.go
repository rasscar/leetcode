func digitSum(s string, k int) string {
	for len(s) > k {
		var nxt []byte
		nxt = make([]byte, 0)
		n := len(s)

		for i := 0; i < n; i += k {
			batas := i + k
			if batas > n {
				batas = n
			}

			jum := 0
			for j := i; j < batas; j++ {
				jum += int(s[j] - '0')
			}

			nxt = append(nxt, strconv.Itoa(jum)...)
		}

		s = string(nxt)
	}
	return s
}
