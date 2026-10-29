func minPatches(nums []int, n int) int {
var target int
target = 1
var hasil int
hasil = 0
var i int
i = 0

for target <= n {
if i < len(nums) && nums[i] <= target {
target += nums[i]
i++
} else {
target += target
hasil++
}
}

return hasil
}
