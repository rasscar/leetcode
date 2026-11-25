
func cek(pf1 []int64, pf2 []int64, p1 []int64, p2 []int64, n int, m int) bool {
	var mod1, mod2 int64
	mod1 = 1000000007
	mod2 = 1000000009

	type pasangan struct {
		h1 int64
		h2 int64
	}

	var mp map[pasangan]int
	mp = make(map[pasangan]int)

	var val1, val2 int64
	val1 = p1[m]
	val2 = p2[m]

	var l, r int
	var h1, h2 int64
	var k pasangan

	for l = 0; l <= n-m; l++ {
		r = l + m

		h1 = (pf1[r] - (pf1[l]*val1)%mod1) % mod1
		if h1 < 0 {
			h1 += mod1
		}

		h2 = (pf2[r] - (pf2[l]*val2)%mod2) % mod2
		if h2 < 0 {
			h2 += mod2
		}

		k.h1 = h1
		k.h2 = h2
		mp[k]++
	}
	for _, c := range mp {
		if c == 1 {
			return true
		}
	}
	return false
}

func smallestUniqueSubarray(nums []int) int {
	var mod1, mod2, b1, b2 int64
	mod1 = 1000000007
	mod2 = 1000000009
	b1 = 1000003
	b2 = 1000033

	var n int
	n = len(nums)

	var pf1, pf2, p1, p2 []int64
	pf1 = make([]int64, n+1)
	pf2 = make([]int64, n+1)
	p1 = make([]int64, n+1)
	p2 = make([]int64, n+1)

	pf1[0] = 0
	pf2[0] = 0
	p1[0] = 1
	p2[0] = 1

	var i int
	var val int64
	for i = 0; i < n; i++ {
		val = int64(nums[i] + 1)
		pf1[i+1] = (pf1[i]*b1 + val) % mod1
		pf2[i+1] = (pf2[i]*b2 + val) % mod2
		p1[i+1] = (p1[i] * b1) % mod1
		p2[i+1] = (p2[i] * b2) % mod2
	}

	var l, r, m, ans int
	l = 1
	r = n
	ans = n

	for l <= r {
		m = l + (r-l)/2
		if cek(pf1, pf2, p1, p2, n, m) {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
