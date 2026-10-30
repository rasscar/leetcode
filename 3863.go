func minOperations(s string) int {
panjang := len(s)
con := true
for i := 0; i < panjang; i++ {
if i > 0 && s[i] < s[i-1] {
con = false
break
}
}
if con {
return 0
}
if panjang == 2 {
return -1
}
minChar := byte('z')
maxChar := byte('a')
for i := 0; i < panjang; i++ {
ch := s[i]
if ch < minChar {
minChar = ch
}
if ch > maxChar {
maxChar = ch
}
}
if minChar == s[0] || maxChar == s[panjang-1] {
return 1
}
for i := 1; i < panjang-1; i++ {
if s[i] == minChar || s[i] == maxChar {
return 2
}
}
return 3
}
