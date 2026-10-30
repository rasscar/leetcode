func numOfWays(nums []int) int {
	mod := int64(1000000007)
	n := len(nums)
	fakt := make([]int64, n+1)
	invers := make([]int64, n+1)
	fakt[0] = 1
	for i := 1; i <= n; i++ {
		fakt[i] = fakt[i-1] * int64(i) % mod
	}
	invers[n] = powMod(fakt[n], mod-2, mod)
	for i := n - 1; i >= 0; i-- {
		invers[i] = invers[i+1] * int64(i+1) % mod
	}
	var dfs func([]int) int64
	dfs = func(arr []int) int64 {
		panjang := len(arr)
		if panjang <= 2 {
			return 1
		}
		kiri := make([]int, 0, panjang-1)
		kanan := make([]int, 0, panjang-1)
		for i := 0; i < panjang; i++ {
			if arr[i] < arr[0] {
				kiri = append(kiri, arr[i])
			} else if arr[i] > arr[0] {
				kanan = append(kanan, arr[i])
			}
		}
		com := fakt[panjang-1] * invers[len(kiri)] % mod * invers[panjang-1-len(kiri)] % mod
		return com * dfs(kiri) % mod * dfs(kanan) % mod
	}
	hasil := dfs(nums)
	hasil = (hasil - 1 + mod) % mod
	return int(hasil)
}

func powMod(base int64, exp int64, mod int64) int64 {
	hasil := int64(1)
	for exp > 0 {
		if exp&1 == 1 {
			hasil = hasil * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return hasil
}
