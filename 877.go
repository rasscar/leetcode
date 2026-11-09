func stoneGame(piles []int) bool {
    sort.Ints(piles)
    c:= len(piles)
    if piles[c-1] + piles[c-3] >= piles[c-2] + piles[c-4]{
        return true
    }else{
        return false
    }
}
