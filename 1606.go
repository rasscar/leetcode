type Job struct {
	at     int
	server int
}

type JobHeap []Job

func (h JobHeap) Len() int { return len(h) }
func (h JobHeap) Less(i, j int) bool {
	if h[i].at != h[j].at {
		return h[i].at < h[j].at
	}
	return h[i].server < h[j].server
}
func (h JobHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *JobHeap) Push(x interface{}) {
	*h = append(*h, x.(Job))
}
func (h *JobHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func busiestServers(k int, arrival []int, load []int) []int {
	var n, i, x, modI, diff, turn, most int
	n = len(arrival)

	var free IntHeap
	free = make(IntHeap, k)
	for i = 0; i < k; i++ {
		free[i] = i
	}
	heap.Init(&free)

	var busy JobHeap
	busy = make(JobHeap, 0)
	heap.Init(&busy)

	var handled []int
	handled = make([]int, k)

	for i = 0; i < n; i++ {
		for len(busy) > 0 && busy[0].at <= arrival[i] {
			x = busy[0].server
			heap.Pop(&busy)

			modI = i % k
			diff = x - modI
			if diff < 0 {
				diff += k
			}
			turn = i + diff

			heap.Push(&free, turn)
		}

		if len(free) == 0 {
			continue
		}

		x = free[0] % k
		heap.Pop(&free)

		var j Job
		j.at = arrival[i] + load[i]
		j.server = x
		heap.Push(&busy, j)
		handled[x]++
	}

	most = 0
	for i = 0; i < k; i++ {
		if handled[i] > most {
			most = handled[i]
		}
	}

	var ans []int
	ans = make([]int, 0)
	for i = 0; i < k; i++ {
		if handled[i] == most {
			ans = append(ans, i)
		}
	}

	return ans
}
