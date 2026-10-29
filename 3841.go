package main

type Pohon struct {
	par, masuk, keluar, urutan []int
	pertama                    [][]int
}

func buatPohon(n int, edges [][]int) *Pohon {
	var i, x, y int
	derajat := make([]int, n)
	for i = 0; i < len(edges); i++ {
		derajat[edges[i][0]]++
		derajat[edges[i][1]]++
	}

	graf := make([][]int, n)
	for i = 0; i < n; i++ {
		graf[i] = make([]int, 0, derajat[i])
	}

	for i = 0; i < len(edges); i++ {
		graf[edges[i][0]] = append(graf[edges[i][0]], edges[i][1])
		graf[edges[i][1]] = append(graf[edges[i][1]], edges[i][0])
	}

	p := &Pohon{
		par:    make([]int, n),
		masuk:  make([]int, n),
		keluar: make([]int, n),
		urutan: make([]int, 0, n),
	}

	tumpukan := []int{0}
	for len(tumpukan) > 0 {
		x = tumpukan[len(tumpukan)-1]
		tumpukan = tumpukan[:len(tumpukan)-1]
		if x < 0 {
			p.keluar[^x] = len(p.urutan)
			continue
		}
		p.masuk[x] = len(p.urutan)
		p.urutan = append(p.urutan, x)
		tumpukan = append(tumpukan, ^x)
		for i = 0; i < len(graf[x]); i++ {
			y = graf[x][i]
			if y != p.par[x] {
				p.par[y] = x
				tumpukan = append(tumpukan, y)
			}
		}
	}

	p.pertama = make([][]int, 1)
	p.pertama[0] = make([]int, n)
	for i = 0; i < len(p.urutan); i++ {
		x = p.urutan[i]
		p.pertama[0][i] = p.masuk[p.par[x]]
	}

	var k, half, ukuran int
	for k = 1; (1 << k) <= n; k++ {
		ukuran = n - (1 << k) + 1
		p.pertama = append(p.pertama, make([]int, ukuran))
		half = 1 << (k - 1)
		for i = 0; i < ukuran; i++ {
			if p.pertama[k-1][i] < p.pertama[k-1][i+half] {
				p.pertama[k][i] = p.pertama[k-1][i]
			} else {
				p.pertama[k][i] = p.pertama[k-1][i+half]
			}
		}
	}

	return p
}

func (p *Pohon) lca(u int, v int) int {
	if u == v {
		return u
	}
	var l, r int
	if p.masuk[u] < p.masuk[v] {
		l = p.masuk[u]
		r = p.masuk[v]
	} else {
		l = p.masuk[v]
		r = p.masuk[u]
	}

	k := hitungBit(r-l) - 1
	var kiri int
	if p.pertama[k][l+1] < p.pertama[k][r-(1<<k)+1] {
		kiri = p.pertama[k][l+1]
	} else {
		kiri = p.pertama[k][r-(1<<k)+1]
	}

	return p.urutan[kiri]
}

type JalurAkar struct {
	pohon         *Pohon
	awal, fenwick []int
}

func buatJalurAkar(t *Pohon, bit []int) *JalurAkar {
	var i, x int
	awal := make([]int, len(bit))
	for i = 0; i < len(bit); i++ {
		awal[i] = bit[i]
	}

	for i = 1; i < len(t.urutan); i++ {
		x = t.urutan[i]
		awal[x] ^= awal[t.par[x]]
	}

	return &JalurAkar{t, awal, make([]int, len(bit)+1)}
}

func (j *JalurAkar) balik(x int, d int) {
	var i, k int
	for k = 0; k < 2; k++ {
		if k == 0 {
			i = j.pohon.masuk[x] + 1
		} else {
			i = j.pohon.keluar[x] + 1
		}
		for i < len(j.fenwick) {
			j.fenwick[i] ^= d
			i += i & -i
		}
	}
}

func (j *JalurAkar) ambil(x int) int {
	var i int
	mask := j.awal[x]
	for i = j.pohon.masuk[x] + 1; i > 0; i -= i & -i {
		mask ^= j.fenwick[i]
	}
	return mask
}

func hitungBit(x int) int {
	var c int
	c = 0
	for x > 0 {
		c++
		x >>= 1
	}
	return c
}

func popCount(x int) int {
	var c int
	c = 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func atoi(s string, l int, r int) int {
	var res, i int
	res = 0
	for i = l; i < r; i++ {
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	var i, x, u, v, spasi, mask int
	pohon := buatPohon(n, edges)
	bit := make([]int, n)
	for i = 0; i < n; i++ {
		bit[i] = 1 << (s[i] - 'a')
	}
	jalur := buatJalurAkar(pohon, bit)

	hasil := make([]bool, 0)
	for i = 0; i < len(queries); i++ {
		q := queries[i]
		if q[0] == 'u' {
			x = atoi(q, 7, len(q)-2)
			var newBit int
			newBit = 1 << (q[len(q)-1] - 'a')
			jalur.balik(x, bit[x]^newBit)
			bit[x] = newBit
		} else {
			for spasi = len(q) - 1; spasi >= 0; spasi-- {
				if q[spasi] == ' ' {
					break
				}
			}
			u = atoi(q, 6, spasi)
			v = atoi(q, spasi+1, len(q))
			mask = jalur.ambil(u) ^ jalur.ambil(v) ^ bit[pohon.lca(u, v)]
			if popCount(mask) <= 1 {
				hasil = append(hasil, true)
			} else {
				hasil = append(hasil, false)
			}
		}
	}
	return hasil
}
