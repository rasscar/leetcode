func findSubtreeSizes(parent []int, s string) []int {
    n:= len(parent)
    nak := make([][]int, n)
    for i:= 1; i<n; i++{
        nak[parent[i]] = append(nak[parent[i]], i)
    }
    uk, last := slices.Repeat([]int{1}, n), slices.Repeat([]int{-1},26)
    var dfs func(i int)
    dfs = func(i int){
        c := s[i] -'a'
        ab := last[c]
        last[c] = i
        for _, child := range nak[i]{
            dfs(child)
        }
        last[c] = ab
        if i > 0 {
            to := parent[i]
            if ab >= 0{
                to = ab
            } 
            uk[to] += uk[i]
        }
    }
    dfs(0)
    return uk
}
