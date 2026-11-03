func minMoves(nums []int) int {
    var max, has int
    for i:=0; i<len(nums); i++{
        if max < nums[i] {
            max = nums[i]
        }
    }
    for i:=0;i<len(nums);i++{
        has += max - nums[i]
    }
    return has
}
