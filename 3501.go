package main

func maxActiveSectionsAfterTrade(s string, qs [][]int) []int {
	n := len(s)
	var rs, re, rl []int
	var rt []byte
	ri := make([]int, n)
	i := 0
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		idx := len(rs)
		rs = append(rs, i)
		re = append(re, j-1)
		rl = append(rl, j-i)
		rt = append(rt, s[i])
		for k := i; k < j; k++ {
			ri[k] = idx
		}
		i = j
	}
	m := len(rs)
	to := 0
	for k := 0; k < n; k++ {
		if s[k] == '1' {
			to++
		}
	}
	const neg = -1 << 30
	gg := make([]int, m)
	for j := 0; j < m; j++ {
		if j-1 >= 0 && j+1 < m && rt[j] == '1' {
			gg[j] = rl[j-1] + rl[j+1]
		} else {
			gg[j] = neg
		}
	}
	ml := 0
	for (1 << uint(ml+1)) <= m {
		ml++
	}
	sp := make([][]int, ml+1)
	sp[0] = append([]int{}, gg...)
	for k := 1; k <= ml; k++ {
		h := 1 << uint(k-1)
		cl := m - (1 << uint(k)) + 1
		if cl < 0 {
			cl = 0
		}
		sp[k] = make([]int, cl)
		for x := 0; x < cl; x++ {
			a := sp[k-1][x]
			b := sp[k-1][x+h]
			if b > a {
				a = b
			}
			sp[k][x] = a
		}
	}
	lg := make([]int, m+2)
	for x := 2; x < len(lg); x++ {
		lg[x] = lg[x/2] + 1
	}
	rm := func(l, r int) int {
		if l > r {
			return neg
		}
		k := lg[r-l+1]
		a := sp[k][l]
		b := sp[k][r-(1<<uint(k))+1]
		if b > a {
			return b
		}
		return a
	}
	ans := make([]int, len(qs))
	for i, qry := range qs {
		l, r := qry[0], qry[1]
		p := ri[l]
		q := ri[r]
		mg := 0
		if p == q || q == p+1 {
			mg = 0
		} else {
			cl := re[p] - l + 1
			cr := r - rs[q] + 1
			if q == p+2 {
				t := p + 1
				if rt[t] == '1' {
					gv := cl + cr
					if gv > mg {
						mg = gv
					}
				}
			} else {
				if rt[p+1] == '1' {
					gv := cl + rl[p+2]
					if gv > mg {
						mg = gv
					}
				}
				if rt[q-1] == '1' {
					gv := rl[q-2] + cr
					if gv > mg {
						mg = gv
					}
				}
				if p+2 <= q-2 {
					gv := rm(p+2, q-2)
					if gv > mg {
						mg = gv
					}
				}
			}
		}
		ans[i] = to + mg
	}
	return ans
}
