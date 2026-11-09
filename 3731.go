func findMissingElements(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    sort.Ints(nums)
    exists := make(map[int]bool)
    for _, v := range nums {
        exists[v] = true
    }
    var missing []int
    for v := nums[0]; v <= nums[len(nums)-1]; v++ {
        if !exists[v] {
            missing = append(missing, v)
        }
    }
    return missing
}
