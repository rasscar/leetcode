func maxDepth(s string) int {
var d,m int
for i:=0;i<len(s);i++{
if s[i] == '(' {
d++
if d > m {
m = d
}
} else if s[i] == ')' {
d--
}
}
return m
}
