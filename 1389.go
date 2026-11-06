func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    reserved := make(map[int]int)
    for _, s := range reservedSeats {
        row, seat := s[0], s[1]
        reserved[row] |= 1 << (seat - 1)
    }

    const (
        leftMask  = 0b11110
        midMask   = 0b1111000
        rightMask = 0b111100000
    )

    ans := (n - len(reserved)) * 2

    for _, mask := range reserved {
        leftOk := mask&leftMask == 0
        midOk := mask&midMask == 0
        rightOk := mask&rightMask == 0

        if leftOk && rightOk {
            ans += 2
        } else if leftOk || midOk || rightOk {
            ans += 1
        }
    }

    return ans
}
