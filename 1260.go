func shiftGrid(grid [][]int, k int) [][]int {
    m := len(grid)
    n := len(grid[0])
    total := m * n
    k = k % total
    if k == 0 {
        return grid
    }

    flat := make([]int, total)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            flat[i*n+j] = grid[i][j]
        }
    }

    newGrid := make([][]int, m)
    for i := 0; i < m; i++ {
        newGrid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            oldIdx := (i*n + j - k + total) % total
            newGrid[i][j] = flat[oldIdx]
        }
    }

    return newGrid
}
