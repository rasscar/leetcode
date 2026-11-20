func findSubstring(s string, words []string) []int {
	n := len(s)
	m := len(words)

	if n == 0 || m == 0 {
		return []int{}
	}

	panjang := len(words[0])
	totalPjg := panjang * m

	if n < totalPjg {
		return []int{}
	}

	var mp map[string]int
	mp = make(map[string]int)

	for i := 0; i < m; i++ {
		mp[words[i]]++
	}

	var hasil []int
	hasil = make([]int, 0)

	for offset := 0; offset < panjang; offset++ {
		var seen map[string]int
		seen = make(map[string]int)
		var size int
		size = 0

		for i := offset; i+panjang <= n; i += panjang {
			var sub string
			sub = s[i : i+panjang]

			var val int
			var ada bool
			val, ada = mp[sub]

			if !ada {
				seen = make(map[string]int)
				size = 0
				continue
			}

			seen[sub]++
			size++

			for seen[sub] > val {
				var first string
				first = s[i-(size-1)*panjang : i-(size-1)*panjang+panjang]
				seen[first]--
				size--
			}

			if size == m {
				hasil = append(hasil, i-(size-1)*panjang)
			}
		}
	}

	return hasil
}
