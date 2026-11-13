func countNonDecreasingSubarrays(nums []int, k int) int64 {
ans := int64(0)
bi := int64(0)
type sm struct {
val, cnt int
}
d := []sm{}
n := len(nums)
j := n - 1

for i := n - 1; i >= 0; i-- {
num := nums[i]
hit := 1

for len(d) > 0 && d[len(d)-1].val < num {
bc := d[len(d)-1]
d = d[:len(d)-1]
hit += bc.cnt
bi += int64(num-bc.val) * int64(bc.cnt)
}

d = append(d, sm{val: num, cnt: hit})

for bi > int64(k) {
front := d[0]
d = d[1:]
bi -= int64(front.val - nums[j])
j--

if front.cnt > 1 {
d = append([]sm{{front.val, front.cnt - 1}}, d...)
}
}

ans += int64(j - i + 1)
}

return ans
}
