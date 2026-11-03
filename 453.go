func minMoves(nums []int) int {
    lo, moves := slices.Min(nums), 0
    for _, x := range nums { moves += x - lo }
    return moves
}
