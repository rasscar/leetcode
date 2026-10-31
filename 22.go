func generateParenthesis(n int) []string {
	hasil := make([]string, 0)
	jalur := make([]byte, 0, 2*n)
	var backtrack func(open int, close int)
	backtrack = func(open int, close int) {
		if len(jalur) == 2*n {
			hasil = append(hasil, string(jalur))
			return
		}
		if open < n {
			jalur = append(jalur, '(')
			backtrack(open+1, close)
			jalur = jalur[:len(jalur)-1]
		}
		if close < open {
			jalur = append(jalur, ')')
			backtrack(open, close+1)
			jalur = jalur[:len(jalur)-1]
		}
	}
	backtrack(0, 0)
	return hasil
}
