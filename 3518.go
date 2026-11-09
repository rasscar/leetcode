func smallestPalindrome(s string, k int) string {
	var y [26]int
	for i := 0; i < len(s); i++ {
		y[s[i]-'a']++
	}

	var half [26]int
	var mid int
	var con bool
	for i := 0; i < 26; i++ {
		half[i] = y[i] / 2
		if y[i]%2 == 1 {
			con = true
			mid = i
		}
	}

	count := func(freq [26]int) int {
		total := 0
		for _, v := range freq {
			total += v
		}
		if total == 0 {
			return 1
		}
		res := 1
		rem := total
		for i := 0; i < 26; i++ {
			if freq[i] == 0 {
				continue
			}
			c := 1
			r := freq[i]
			if r > rem-r {
				r = rem - r
			}
			for j := 1; j <= r; j++ {
				c = c * (rem - r + j) / j
				if c > 1000001 {
					c = 1000001
					break
				}
			}
			if res > 1000001/c {
				return 1000001
			}
			res *= c
			if res > 1000001 {
				return 1000001
			}
			rem -= freq[i]
		}
		return res
	}

	total := count(half)
	if k > total {
		return ""
	}

	n := len(s) / 2
	a := make([]byte, n)
	rem := half
	for pos := 0; pos < n; pos++ {
		for c := 0; c < 26; c++ {
			if rem[c] == 0 {
				continue
			}
			rem[c]--
			cnt := count(rem)
			rem[c]++
			if k > cnt {
				k -= cnt
				continue
			}
			a[pos] = byte('a' + c)
			rem[c]--
			break
		}
	}

	hasil := make([]byte, 0, len(s))
	hasil = append(hasil, a...)
	if con {
		hasil = append(hasil, byte('a'+mid))
	}
	for i := n - 1; i >= 0; i-- {
		hasil = append(hasil, a[i])
	}
	return string(hasil)

}
