func canJump(nums []int) bool {
n := len(nums)
if n == 1 {
return true
}
batasSekarang := 0
for i := 0; i <= batasSekarang && i < n; i++ {
temp := i + nums[i]
if temp > batasSekarang {
batasSekarang = temp
}
if batasSekarang >= n-1 {
return true
}
}
return false
}
