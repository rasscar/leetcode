func countServers(grid [][]int) int {
    rows, cols := make([]int, len(grid)), make([]int, len(grid[0]))  
    for i, row := range grid {
        for j, server := range row {
            rows[i] += server
            cols[j] += server
        }
    }

    ans := 0
    for i, row := range grid {
        for j, server := range row {
            if server == 1 && rows[i]+cols[j] > 2 {  
                ans++
            }
        }
    }
    return ans
}
