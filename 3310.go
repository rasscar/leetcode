func remainingMethods(n int, k int, invocations [][]int) []int {
    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        adj[i] = []int{}
    }
    for _, e := range invocations {
        a, b := e[0], e[1]
        adj[a] = append(adj[a], b)
    }

    suspicious := make([]bool, n)
    queue := []int{k}
    suspicious[k] = true
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        for _, v := range adj[u] {
            if !suspicious[v] {
                suspicious[v] = true
                queue = append(queue, v)
            }
        }
    }

    for _, e := range invocations {
        a, b := e[0], e[1]
        if !suspicious[a] && suspicious[b] {
            semua := make([]int, n)
            for i := 0; i < n; i++ {
                semua[i] = i
            }
            return semua
        }
    }

    hasil := []int{}
    for i := 0; i < n; i++ {
        if !suspicious[i] {
            hasil = append(hasil, i)
        }
    }
    return hasil
}
