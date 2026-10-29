func restoreIpAddresses(s string) []string {
var n int
n = len(s)
var hasil []string
hasil = make([]string, 0)

var cekValid func(temp string) bool
cekValid = func(temp string) bool {
if len(temp) > 3 || len(temp) == 0 {
return false
}
if len(temp) > 1 && temp[0] == '0' {
return false
}
var nilai, i int
nilai = 0
for i = 0; i < len(temp); i++ {
nilai = nilai*10 + int(temp[i]-'0')
}
if nilai > 255 {
return false
}
return true
}

var dfs func(ind int, dots int, current string)
dfs = func(ind int, dots int, current string) {
if dots == 3 {
if cekValid(s[ind:]) {
hasil = append(hasil, current+s[ind:])
}
return
}

var batas, i int
batas = ind + 3
if batas > n {
batas = n
}

for i = ind; i < batas; i++ {
var sub string
sub = s[ind : i+1]
if cekValid(sub) {
var temp string
temp = current + sub + "."
dfs(i+1, dots+1, temp)
}
}
}

dfs(0, 0, "")
return hasil
}
