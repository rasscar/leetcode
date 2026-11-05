func reverseDegree(s string) int {
	totalSum := 0
	for i, char := range s {
		revAlphabetIndex := 'z' - char + 1
		positionInString := i + 1
		totalSum += positionInString * int(revAlphabetIndex)
	}
	return totalSum
}
