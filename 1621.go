const LIM = 1999
const MOD = 1000000007
var fact [LIM]int
var inv [LIM]int

func init() {
	fact[0] = 1

	for i := 1; i < LIM; i++ {
		fact[i] = fact[i-1] * i % MOD
	}

	inv[LIM-1] = pow(fact[LIM-1], MOD-2)

	for i := LIM - 1; i > 0; i-- {
		inv[i-1] = inv[i] * i % MOD
	}
}

func pow(x, n int) int {
	ans := 1

	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			ans = ans * x % MOD
		}

		x = x * x % MOD
	}

	return ans
}

func comb(n, m int) int {
	return fact[n] * inv[m] % MOD * inv[n-m] % MOD
}

func numberOfSets(n, k int) int {
	return comb(n+k-1, k*2)
}
