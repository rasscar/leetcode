func countIslands(grid [][]int, k int) int {
var m, n int
m = len(grid)
n = len(grid[0])
var hasil int
hasil = 0

var dfs func(i int, j int) int64
dfs = func(i int, j int) int64 {
if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
return 0
}
var jumlah int64
jumlah = int64(grid[i][j])
grid[i][j] = 0
jumlah += dfs(i+1, j)
jumlah += dfs(i-1, j)
jumlah += dfs(i, j+1)
jumlah += dfs(i, j-1)
return jumlah
}

var i, j int
for i = 0; i < m; i++ {
for j = 0; j < n; j++ {
if grid[i][j] > 0 {
var skr int64
skr = dfs(i, j)
if skr%int64(k) == 0 {
hasil++
}
}
}
}
return hasil
}
