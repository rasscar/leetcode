func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)

	type point struct {
		x, y int
	}

	var ones1, ones2 []point

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img1[i][j] == 1 {
				ones1 = append(ones1, point{i, j})
			}
			if img2[i][j] == 1 {
				ones2 = append(ones2, point{i, j})
			}
		}
	}

	shiftCount := make(map[point]int)
	maxOverlap := 0

	for _, p1 := range ones1 {
		for _, p2 := range ones2 {
			dx := p2.x - p1.x
			dy := p2.y - p1.y
			shift := point{dx, dy}

			shiftCount[shift]++

			if shiftCount[shift] > maxOverlap {
				maxOverlap = shiftCount[shift]
			}
		}
	}

	return maxOverlap
}
