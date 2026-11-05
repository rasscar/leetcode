type Interval struct {
	l, r, w, idx int
}

type State struct {
	score   int
	indices []int
}

func isBetter(a, b State) bool {
	if a.score != b.score {
		return a.score > b.score
	}
	aIdx := make([]int, len(a.indices))
	copy(aIdx, a.indices)
	sort.Ints(aIdx)
	bIdx := make([]int, len(b.indices))
	copy(bIdx, b.indices)
	sort.Ints(bIdx)

	for i := 0; i < len(aIdx) && i < len(bIdx); i++ {
		if aIdx[i] != bIdx[i] {
			return aIdx[i] < bIdx[i]
		}
	}
	return len(aIdx) < len(bIdx)
}

func maximumWeight(intervals [][]int) []int {
	n := len(intervals)
	items := make([]Interval, n)
	for i, v := range intervals {
		items[i] = Interval{l: v[0], r: v[1], w: v[2], idx: i}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].r != items[j].r {
			return items[i].r < items[j].r
		}
		if items[i].l != items[j].l {
			return items[i].l < items[j].l
		}
		return items[i].idx < items[j].idx
	})

	prev := make([]int, n)
	for i := 0; i < n; i++ {
		j := sort.Search(n, func(k int) bool {
			return items[k].r >= items[i].l
		})
		prev[i] = j - 1
	}

	dp := make([][]State, 5)
	for k := 0; k <= 4; k++ {
		dp[k] = make([]State, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = State{score: 0, indices: []int{}}
	}

	for k := 1; k <= 4; k++ {
		for i := 0; i < n; i++ {
			if i > 0 {
				dp[k][i] = dp[k][i-1]
			} else {
				dp[k][i] = State{score: 0, indices: []int{}}
			}

			var candidate State
			if prev[i] != -1 {
				prevState := dp[k-1][prev[i]]
				candidate.score = items[i].w + prevState.score
				candidate.indices = append([]int{}, prevState.indices...)
				candidate.indices = append(candidate.indices, items[i].idx)
			} else {
				candidate.score = items[i].w
				candidate.indices = []int{items[i].idx}
			}

			if isBetter(candidate, dp[k][i]) {
				dp[k][i] = candidate
			}
		}
	}

	res := dp[4][n-1].indices
	sort.Ints(res)
	return res
}
