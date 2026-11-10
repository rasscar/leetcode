func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	hasil := 0

	for i < j {
		luas := (j - i) * min(height[i], height[j])
		if luas > hasil {
			hasil = luas
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return hasil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
