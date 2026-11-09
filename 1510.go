func winnerSquareGame(n int) bool {
    menang := make([]bool, n+1)
    for i := 1; i <= n; i++ {
        for s := 1; s*s <= i; s++ {
            if !menang[i-s*s] {
                menang[i] = true
                break
            }
        }
    }
    return menang[n]
}
