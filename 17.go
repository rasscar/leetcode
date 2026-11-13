func letterCombinations(digits string) []string {
var hasil []string
hasil = make([]string, 0)

if len(digits) == 0 {
return hasil
}

var mp [256]string
mp['2'] = "abc"
mp['3'] = "def"
mp['4'] = "ghi"
mp['5'] = "jkl"
mp['6'] = "mno"
mp['7'] = "pqrs"
mp['8'] = "tuv"
mp['9'] = "wxyz"

var dfs func(idx int, jalur string)
dfs = func(idx int, jalur string) {
if idx == len(digits) {
hasil = append(hasil, jalur)
return
}

var huruf string
huruf = mp[digits[idx]]

var i int
for i = 0; i < len(huruf); i++ {
dfs(idx+1, jalur+string(huruf[i]))
}
}

dfs(0, "")

return hasil
}

