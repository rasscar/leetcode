func largestInteger(nums []int, k int) int {
    n := len(nums)

    pos := make(map[int][]int)
    for i, x := range nums {
        pos[x] = append(pos[x], i)
    }

    totalSubarrays := n - k + 1
    ans := -1
    found := false

    for x, p := range pos {
        notContaining := 0

        if p[0]-k+1 > 0 {
            notContaining += p[0] - k + 1
        }

        for i := 0; i+1 < len(p); i++ {
            gap := p[i+1] - p[i] - k
            if gap > 0 {
                notContaining += gap
            }
        }

        if n-p[len(p)-1]-k > 0 {
            notContaining += n - p[len(p)-1] - k
        }

        containing := totalSubarrays - notContaining

        if containing == 1 {
            if !found || x > ans {
                ans = x
                found = true
            }
        }
    }

    if !found {
        return -1
    }
    return ans
}
