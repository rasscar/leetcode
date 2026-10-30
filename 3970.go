type Edge struct {
	v, w int
}

type State struct {
	jarak  int64
	simpul int
	run    int
}

func push(h []State, s State) []State {
	h = append(h, s)
	var i, p int
	i = len(h) - 1
	for i > 0 {
		p = (i - 1) / 2
		if h[p].jarak <= h[i].jarak {
			break
		}
		h[p], h[i] = h[i], h[p]
		i = p
	}
	return h
}

func pop(h []State) ([]State, State) {
	var n int
	n = len(h)
	if n == 0 {
		return h, State{}
	}
	var top State
	top = h[0]
	h[0] = h[n-1]
	h = h[:n-1]
	var i, l, r, m int
	i = 0
	for {
		l = 2*i + 1
		r = 2*i + 2
		m = i
		if l < len(h) && h[l].jarak < h[m].jarak {
			m = l
		}
		if r < len(h) && h[r].jarak < h[m].jarak {
			m = r
		}
		if m == i {
			break
		}
		h[i], h[m] = h[m], h[i]
		i = m
	}
	return h, top
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
	var adj [][]Edge
	adj = make([][]Edge, n)
	var i, j int
	for i = 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], Edge{edges[i][1], edges[i][2]})
	}

	var dist [][]int64
	dist = make([][]int64, n)
	for i = 0; i < n; i++ {
		dist[i] = make([]int64, k+1)
		for j = 0; j <= k; j++ {
			dist[i][j] = 1 << 60
		}
	}

	dist[0][1] = 0
	var antrian []State
	antrian = make([]State, 0)
	antrian = push(antrian, State{0, 0, 1})

	for len(antrian) > 0 {
		var skr State
		antrian, skr = pop(antrian)

		if skr.jarak > dist[skr.simpul][skr.run] {
			continue
		}

		for i = 0; i < len(adj[skr.simpul]); i++ {
			var v, w int
			v = adj[skr.simpul][i].v
			w = adj[skr.simpul][i].w

			var runBaru int
			runBaru = 1
			if labels[v] == labels[skr.simpul] {
				runBaru = skr.run + 1
			}
			if runBaru > k {
				continue
			}

			var jarakBaru int64
			jarakBaru = skr.jarak + int64(w)
			if jarakBaru < dist[v][runBaru] {
				dist[v][runBaru] = jarakBaru
				antrian = push(antrian, State{jarakBaru, v, runBaru})
			}
		}
	}

	var hasil int64
	hasil = 1 << 60
	for i = 0; i <= k; i++ {
		if dist[n-1][i] < hasil {
			hasil = dist[n-1][i]
		}
	}

	if hasil == 1 << 60 {
		return -1
	}
	return int(hasil)
}
