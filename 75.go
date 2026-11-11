func sortColors(nums []int) {
x := 0
y := 0
z := 0
for i := 0; i < len(nums); i++ {
if nums[i] == 0 {
x++
} else if nums[i] == 1 {
y++
} else {
z++
}
}
for i := 0; i < len(nums); i++ {
if x > 0 {
nums[i] = 0
x--
} else if y > 0 {
nums[i] = 1
y--
} else {
nums[i] = 2
z--
}
}
}
