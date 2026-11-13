func intToRoman(num int) string {
	var vals [13]int
	var syms [13]string
	var res []byte
	var i int

	vals = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms = [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res = make([]byte, 0)

	for i = 0; i < 13; i++ {
		for num >= vals[i] {
			res = append(res, syms[i]...)
			num -= vals[i]
		}
	}

	return string(res)
}
