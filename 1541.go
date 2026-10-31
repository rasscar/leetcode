func minInsertions(s string) int {
	var count, depth int32
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			depth++
		} else {
			depth--
			if depth < 0 {
				count++
				depth++
			}
			i++
			if i == len(s) {
				count++
			} else if s[i] == '(' {
				count++
				depth++
			}
		}
		i++
	}
	if depth > 0 {
		count += 2 * depth
	}
	return int(count)
}
