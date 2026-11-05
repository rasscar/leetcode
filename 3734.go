func lexPalindromicPermutation(s string, target string) string {
	n := len(s)
	cnt := [26]int{}
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}

	odd := 0
	mid := byte(0)
	for c := 0; c < 26; c++ {
		if cnt[c]%2 == 1 {
			odd++
			mid = byte('a' + c)
		}
	}

	if odd > 1 || (n%2 == 1 && odd != 1) || (n%2 == 0 && odd != 0) {
		return ""
	}

	if n%2 == 0 {
		m := n / 2
		half := [26]int{}
		for c := 0; c < 26; c++ {
			half[c] = cnt[c] / 2
		}

		h0 := target[:m]
		if isValid(h0, half) {
			cand := buildPalindrome(h0, n)
			if cand > target {
				return cand
			}
		}

		nh := nextPermutation(h0, half, m)
		if nh == "" {
			return ""
		}
		return buildPalindrome(nh, n)
	}

	leftLen := n / 2
	pair := [26]int{}
	for c := 0; c < 26; c++ {
		pair[c] = cnt[c] / 2
	}

	L0 := target[:leftLen]
	mid0 := target[leftLen]

	if mid0 == mid && isValid(L0, pair) {
		h0 := L0 + string(mid)
		cand := buildPalindrome(h0, n)
		if cand > target {
			return cand
		}
	}

	if mid > mid0 && isValid(L0, pair) {
		return buildPalindrome(L0+string(mid), n)
	}

	nl := nextPermutation(L0, pair, leftLen)
	if nl == "" {
		return ""
	}
	return buildPalindrome(nl+string(mid), n)
}

func isValid(prefix string, counts [26]int) bool {
	need := [26]int{}
	for i := 0; i < len(prefix); i++ {
		idx := int(prefix[i] - 'a')
		need[idx]++
		if need[idx] > counts[idx] {
			return false
		}
	}
	return true
}

func buildPalindrome(h string, n int) string {
	res := []byte(h)
	if n%2 == 0 {
		for i := len(h) - 1; i >= 0; i-- {
			res = append(res, h[i])
		}
	} else {
		for i := len(h) - 2; i >= 0; i-- {
			res = append(res, h[i])
		}
	}
	return string(res)
}

func nextPermutation(prev string, counts [26]int, length int) string {
	if length == 0 {
		return ""
	}

	for i := length - 1; i >= 0; i-- {
		prefix := prev[:i]
		if !isValid(prefix, counts) {
			continue
		}

		used := [26]int{}
		for j := 0; j < i; j++ {
			used[prev[j]-'a']++
		}

		rem := counts
		for c := 0; c < 26; c++ {
			rem[c] -= used[c]
		}

		cur := int(prev[i] - 'a')
		c := cur + 1
		for c < 26 && rem[c] == 0 {
			c++
		}
		if c == 26 {
			continue
		}

		res := make([]byte, 0, length)
		res = append(res, prefix...)
		res = append(res, byte('a'+c))
		rem[c]--

		for ch := 0; ch < 26; ch++ {
			for k := 0; k < rem[ch]; k++ {
				res = append(res, byte('a'+ch))
			}
		}

		return string(res)
	}

	return ""
}
