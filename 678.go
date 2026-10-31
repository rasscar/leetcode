func checkValidString(s string) bool {
l := 0
h := 0
for i := 0; i < len(s); i++ {
c := s[i]
if c == '(' {
l++
} else {
l--
}
if c != ')' {
h++
} else {
h--
}
if h < 0 {
return false
}
if l < 0 {
l = 0
}
}
return l == 0
}
