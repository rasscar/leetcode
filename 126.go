func findLadders(bw string, ew string, wl []string) [][]string {
	var ans [][]string
	ans = make([][]string, 0)

	var words map[string]bool
	words = make(map[string]bool)
	var i int
	for i = 0; i < len(wl); i++ {
		words[wl[i]] = true
	}

	if !words[ew] {
		return ans
	}

	delete(words, bw)

	var q []string
	q = make([]string, 0)
	q = append(q, bw)

	var mpp map[string]int
	mpp = make(map[string]int)
	mpp[bw] = 1

	var steps, j int
	steps = 1

	for len(q) > 0 {
		var word string
		word = q[0]
		q = q[1:]
		steps = mpp[word]

		var b []byte
		b = []byte(word)
		
		for j = 0; j < len(b); j++ {
			var original byte
			original = b[j]
			var ch byte
			for ch = 'a'; ch <= 'z'; ch++ {
				b[j] = ch
				var nw string
				nw = string(b)
				if words[nw] {
					mpp[nw] = steps + 1
					q = append(q, nw)
					delete(words, nw)
				}
			}
			b[j] = original
		}
	}

	if _, ada := mpp[ew]; !ada {
		return ans
	}

	var seq []string
	seq = make([]string, 0)
	seq = append(seq, ew)

	var dfs func(word string)
	dfs = func(word string) {
		if word == bw {
			var salin []string
			salin = make([]string, len(seq))
			var x int
			for x = 0; x < len(seq); x++ {
				salin[x] = seq[len(seq)-1-x]
			}
			ans = append(ans, salin)
			return
		}

		var val int
		val = mpp[word]

		var b []byte
		b = []byte(word)
		var x int
		
		for x = 0; x < len(b); x++ {
			var original byte
			original = b[x]
			var ch byte
			for ch = 'a'; ch <= 'z'; ch++ {
				b[x] = ch
				var nw string
				nw = string(b)
				
				if v, ada := mpp[nw]; ada && v == val-1 {
					seq = append(seq, nw)
					dfs(nw)
					seq = seq[:len(seq)-1]
				}
			}
			b[x] = original
		}
	}

	dfs(ew)
	return ans
}
