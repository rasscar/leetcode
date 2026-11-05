func stoneGameVIII(stones []int) int {
    n := len(stones)
    pref := make([]int, n)
    pref[0] = stones[0]
    for i := 1; i < n; i++ {
        pref[i] = pref[i-1] + stones[i]
    }

    // dp merepresentasikan selisih skor maksimal yang bisa didapat pemain saat ini
    // Base case: jika mengambil seluruh batu yang tersisa (indeks terakhir n-1)
    dp := pref[n-1]

    // Transisi mundur dari n-2 hingga indeks 1 (karena x > 1, minimal ambil sampai indeks 1)
    for i := n - 2; i >= 1; i-- {
        if pref[i]-dp > dp {
            dp = pref[i] - dp
        }
    }

    return dp
}
