func minimumPushes(word string) int {
	var x [123]int
	var has int
	var z [27]int
	for i := 0; i < len(word); i++ {
		x[word[i]]++
	}
	for i := 97; i <= 122; i++ {
		z[x[i]] += 1
	}
	bed := 1
	processed := 0
	for i := 26; i >= 1; i-- {
		count := z[i]
		for j := 0; j < count; j++ {
			has += i * bed
			processed++
			if processed%8 == 0 {
				bed++
			}
		}
	}
	return has
}
