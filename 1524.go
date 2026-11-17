func numOfSubarrays(arr []int) int {
	var n, i, jum, x, y int
	n = len(arr)
	jum = 0
	x = 0

	for i = 0; i < n; i++ {
		jum += arr[i]
		if jum%2 != 0 {
			x++
		}
	}

	y = n + 1 - x
	var hasil int64
	hasil = int64(x) * int64(y) % 1000000007
	return int(hasil)
}
