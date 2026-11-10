func nextPermutation(nums []int) {
n := len(nums)
if n <= 1 {
return
}
puncak := n - 1
for puncak > 0 && nums[puncak-1] >= nums[puncak] {
puncak--
}
if puncak != 0 {
besar := n - 1
for ; besar >= puncak; besar-- {
if nums[puncak-1] < nums[besar] {
nums[besar], nums[puncak-1] = nums[puncak-1], nums[besar]
break
}
}
}
kiri := puncak
kanan := n - 1
for kiri < kanan {
nums[kiri], nums[kanan] = nums[kanan], nums[kiri]
kiri++
kanan--
}
}
