func missingMultiple(nums []int, k int) int {
    var x,p int
    var y bool
    p = len(nums)
    y = true
    sort.Ints(nums)
    for y{
        y = false
        x = x + k
        for j:=0; j<p; j++{
            if x == nums[j]{
                y = true
                break
            }
        }
    }
    return x
}
