func minMoves(classroom []string, energy int) int {
	m := len(classroom)
	n := len(classroom[0])

	var sr, sc int
	litter := make([][2]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch classroom[i][j] {
			case 'S':
				sr, sc = i, j
			case 'L':
				litter = append(litter, [2]int{i, j})
			}
		}
	}

	k := len(litter)
	if k == 0 {
		return 0
	}

	idxMap := make(map[[2]int]int)
	for idx, pos := range litter {
		idxMap[pos] = idx
	}

	allMask := (1 << k) - 1
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	type state struct {
		r, c int
		mask int
		e    int
		dist int
	}

	visited := make(map[int]int)
	startKey := ((sr*n + sc) << k) | 0
	visited[startKey] = energy

	queue := []state{{sr, sc, 0, energy, 0}}

	for head := 0; head < len(queue); head++ {
		cur := queue[head]

		if cur.mask == allMask {
			return cur.dist
		}

		if cur.e == 0 {
			continue
		}

		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]

			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			cell := classroom[nr][nc]
			if cell == 'X' {
				continue
			}

			e2 := cur.e - 1
			newMask := cur.mask

			if cell == 'L' {
				if bit, ok := idxMap[[2]int{nr, nc}]; ok {
					newMask |= 1 << bit
				}
			}

			if cell == 'R' {
				e2 = energy
			}

			if newMask == allMask {
				return cur.dist + 1
			}

			if e2 == 0 {
				continue
			}

			key := ((nr*n + nc) << k) | newMask
			if prev, ok := visited[key]; ok && prev >= e2 {
				continue
			}
			visited[key] = e2

			queue = append(queue, state{nr, nc, newMask, e2, cur.dist + 1})
		}
	}

	return -1
}
