func lexicographicallySmallestArray(nums []int, limit int) []int {
    n := len(nums)
    idx := make([]int, n)
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool {
        return nums[idx[i]] < nums[idx[j]]
    })
    ans := make([]int, n)
    for i := 0; i < n; {
        j := i
        for j+1 < n && nums[idx[j+1]]-nums[idx[j]] <= limit {
            j++
        }
        positions := make([]int, j-i+1)
        copy(positions, idx[i:j+1])
        sort.Ints(positions)
        for k, pos := range positions {
            ans[pos] = nums[idx[i+k]]
        }
        i = j + 1
    }
    return ans
}
