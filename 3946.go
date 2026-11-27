func maximumSaleItems(items [][]int, budget int) int {
	var n, top, cheapest, i, f, p, k, money, gain, val int
	n = len(items)
	if n == 0 {
		return 0
	}

	top = 0
	cheapest = 1 << 30

	for i = 0; i < n; i++ {
		f = items[i][0]
		p = items[i][1]
		if f > top {
			top = f
		}
		if p < cheapest {
			cheapest = p
		}
	}

	var have, gains []int
	have = make([]int, top+1)
	gains = make([]int, top+1)

	for i = 0; i < n; i++ {
		have[items[i][0]]++
	}

	for f = 1; f <= top; f++ {
		for k = f; k <= top; k += f {
			gains[f] += have[k]
		}
	}

	var best []int
	best = make([]int, budget+1)
	for money = 0; money <= budget; money++ {
		best[money] = money / cheapest
	}

	for i = 0; i < n; i++ {
		f = items[i][0]
		p = items[i][1]
		gain = gains[f]

		if gain > p/cheapest {
			for money = budget; money >= p; money-- {
				val = best[money-p] + gain
				if val > best[money] {
					best[money] = val
				}
			}
		}
	}

	return best[budget]
}
