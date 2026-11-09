func smallestPalindrome(s string) string {
	var y [123]int
	var con bool
	a := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		x := s[i]
		y[x] += 1
	}
	j := 0
	posisitengah := len(s) - 1
	var tengah int
	for i := 97; i <= 122; i++ {
		b := y[i]
		if b%2 == 1 {
			con = true
			tengah = i
		}
		setengah := b / 2
		for k := 0; k < setengah; k++ {
			a[j] = byte(i)
			a[posisitengah] = byte(i)
			j++
			posisitengah--
		}
	}
	if con {
		a[len(s)/2] = byte(tengah)
	}
	return string(a)
}
