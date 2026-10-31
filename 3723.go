var nines = strings.Repeat("9", 200_000)
var zeros = strings.Repeat("0", 200_000)
var rem = []string{"", "1", "2", "3", "4", "5", "6", "7", "8"}

func maxSumOfSquares(n int, s int) string {
if s > 9*n {
return ""
}
mid := rem[s%9]
x := s / 9
y := n - x - len(mid)
return nines[:x] + mid + zeros[:y]
}

