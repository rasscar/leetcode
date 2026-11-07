import "container/heap"

type taskServer struct {
	freeAt int64
	weight int
	index  int
}

type taskServerHeap []taskServer

func (h taskServerHeap) Len() int { return len(h) }

func (h taskServerHeap) Less(i, j int) bool {
	if h[i].freeAt != h[j].freeAt {
		return h[i].freeAt < h[j].freeAt
	}
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}

func (h taskServerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *taskServerHeap) Push(x interface{}) {
	*h = append(*h, x.(taskServer))
}

func (h *taskServerHeap) Pop() interface{} {
	old := *h
	last := old[len(old)-1]
	*h = old[:len(old)-1]
	return last
}

func assignTasks(servers []int, tasks []int) []int {
	free := make(taskServerHeap, len(servers))
	for i, weight := range servers {
		free[i] = taskServer{weight: weight, index: i}
	}
	heap.Init(&free)

	busy := make(taskServerHeap, 0, len(servers))
	ans := make([]int, len(tasks))
	var now int64

	for j, duration := range tasks {
		if now < int64(j) {
			now = int64(j)
		}
		if len(free) == 0 && busy[0].freeAt > now {
			now = busy[0].freeAt
		}

		for len(busy) > 0 && busy[0].freeAt <= now {
			server := heap.Pop(&busy).(taskServer)
			server.freeAt = 0
			heap.Push(&free, server)
		}

		server := heap.Pop(&free).(taskServer)
		ans[j] = server.index
		server.freeAt = now + int64(duration)
		heap.Push(&busy, server)
	}

	return ans
}
