func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	var ans int64
	var m1, sisa1, sisa2 int
	var c1, c2 int

	if need1 < need2 {
		m1 = need1
	} else {
		m1 = need2
	}

	if cost1+cost2 < costBoth {
		ans = int64(cost1)*int64(need1) + int64(cost2)*int64(need2)
		return ans
	}

	sisa1 = need1 - m1
	sisa2 = need2 - m1

	c1 = cost1
	if costBoth < c1 {
		c1 = costBoth
	}

	c2 = cost2
	if costBoth < c2 {
		c2 = costBoth
	}

	ans = int64(costBoth)*int64(m1) + int64(c1)*int64(sisa1) + int64(c2)*int64(sisa2)
	return ans
}
