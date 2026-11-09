func longestSubsequence(nums []int) int {
    var x, has int
    nonZero := 0

    for true {
        for i := 0; i < len(nums); i++ {
            has ^= nums[i]
            if nums[i] != 0 {
                nonZero++
            }
        }
        break
    }

    if has != 0 {
        x = len(nums)
    } else if nonZero > 0 {
        x = len(nums) - 1
    } else {
        x = 0
    }
    return x
}
