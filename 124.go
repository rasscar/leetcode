func maxPathSum(root *TreeNode) int {
var hasil int
hasil = -1 << 30

var dfs func(node *TreeNode) int
dfs = func(node *TreeNode) int {
if node == nil {
return 0
}

var kiri, kanan int
kiri = dfs(node.Left)
if kiri < 0 {
kiri = 0
}

kanan = dfs(node.Right)
if kanan < 0 {
kanan = 0
}

var skr int
skr = node.Val + kiri + kanan
if skr > hasil {
hasil = skr
}

var terbaik int
if kiri > kanan {
terbaik = node.Val + kiri
} else {
terbaik = node.Val + kanan
}

return terbaik
}

dfs(root)
return hasil
}

