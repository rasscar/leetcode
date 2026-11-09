func uniqueXorTriplets(nums []int) int {
	uniq := make(map[int]bool)
	maxVal := 0
	for _, v := range nums {
		uniq[v] = true
		if v > maxVal {
			maxVal = v
		}
	}
	if len(uniq) == 0 {
		return 0
	}
	K := 1
	for K <= maxVal {
		K <<= 1
	}
	A := make([]int64, K)
	for v := range uniq {
		A[v] = 1
	}
	fwht(A)
	for i := range A {
		A[i] = A[i] * A[i] * A[i]
	}
	fwht(A)
	for i := range A {
		A[i] /= int64(K)
	}
	count := 0
	for _, v := range A {
		if v > 0 {
			count++
		}
	}
	return count
}

func fwht(a []int64) {
	n := len(a)
	for length := 1; length < n; length <<= 1 {
		for i := 0; i < n; i += length << 1 {
			for j := 0; j < length; j++ {
				u := a[i+j]
				v := a[i+j+length]
				a[i+j] = u + v
				a[i+j+length] = u - v
			}
		}
	}
}
