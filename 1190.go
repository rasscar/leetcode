func reverseParentheses(s string) string {
	parpos := []int{}
	ans := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			parpos = append(parpos, len(ans))
		} else if s[i] == ')' {
			j := parpos[len(parpos)-1]
			parpos = parpos[:len(parpos)-1]

			for l, r := j, len(ans)-1; l < r; l, r = l+1, r-1 {
				ans[l], ans[r] = ans[r], ans[l]
			}
		} else {
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}
