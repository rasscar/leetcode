func myAtoi(s string) int {
	var n, i, tanda, digit int
	var hasil, val, maxV, minV int64

	n = len(s)
	i = 0
	tanda = 1
	hasil = 0
	maxV = 1<<31 - 1
	minV = -1 << 31

	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		tanda = -1
		i++
	}

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit = int(s[i] - '0')
		hasil = hasil*10 + int64(digit)
		val = hasil * int64(tanda)

		if val <= minV {
			return int(minV)
		}
		if val >= maxV {
			return int(maxV)
		}

		i++
	}

	return int(val)
}
