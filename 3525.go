func resultArray(nums []int, k int, queries [][]int) []int {
	type node struct {
		product int
		freq    []int
	}

	n := len(nums)
	tree := make([]node, 4*n)

	for i := range tree {
		tree[i] = node{
			product: 1,
			freq:    make([]int, k),
		}
	}

	merge := func(left, right node) node {
		res := node{
			product: 1,
			freq:    make([]int, k),
		}

		res.product = (left.product * right.product) % k
		copy(res.freq, left.freq)

		for r := 0; r < k; r++ {
			if right.freq[r] == 0 {
				continue
			}

			nr := (left.product * r) % k
			res.freq[nr] += right.freq[r]
		}

		return res
	}

	var build func(v, left, right int)
	build = func(v, left, right int) {
		if left == right {
			tree[v].product = nums[left] % k
			tree[v].freq[tree[v].product] = 1
			return
		}

		mid := (left + right) / 2
		build(v*2, left, mid)
		build(v*2+1, mid+1, right)
		tree[v] = merge(tree[v*2], tree[v*2+1])
	}

	var update func(v, left, right, pos, value int)
	update = func(v, left, right, pos, value int) {
		if left == right {
			tree[v].product = value % k

			for i := range tree[v].freq {
				tree[v].freq[i] = 0
			}

			tree[v].freq[tree[v].product] = 1
			return
		}

		mid := (left + right) / 2
		if pos <= mid {
			update(v*2, left, mid, pos, value)
		} else {
			update(v*2+1, mid+1, right, pos, value)
		}

		tree[v] = merge(tree[v*2], tree[v*2+1])
	}

	var query func(v, left, right, ql, qr int) node
	query = func(v, left, right, ql, qr int) node {
		if ql > qr {
			return node{
				product: 1,
				freq:    make([]int, k),
			}
		}

		if ql == left && qr == right {
			return tree[v]
		}

		mid := (left + right) / 2

		return merge(
			query(v*2, left, mid, ql, minInt(qr, mid)),
			query(v*2+1, mid+1, right, maxInt(ql, mid+1), qr),
		)
	}

	build(1, 0, n-1)

	ans := make([]int, 0, len(queries))

	for _, q := range queries {
		idx, val, start, x := q[0], q[1], q[2], q[3]

		update(1, 0, n-1, idx, val)

		res := query(1, 0, n-1, start, n-1)
		ans = append(ans, res.freq[x])
	}

	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
