func reverse(x int) int {
const maxInt32 = 1<<31 - 1
const minInt32 = -1 << 31
var rev int
rev = 0
for x != 0 {
var ld int
ld = x % 10
if rev > maxInt32/10 || (rev == maxInt32/10 && ld > 7) {
return 0
}
if rev < minInt32/10 || (rev == minInt32/10 && ld < -8) {
return 0
}
rev = rev*10 + ld
x = x / 10
}
return rev
}

