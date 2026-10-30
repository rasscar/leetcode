func targetIndices(nums []int, target int) []int {
    var has []int
    var y int
    y = len(nums)
    sort.Ints(nums)
    for i:=0; i<y; i++{
        if nums[i] == target{
            has = append(has, i)
        }
    }
    return has
}
