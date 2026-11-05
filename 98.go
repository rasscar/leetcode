func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, l int64, r int64) bool
	dfs = func(node *TreeNode, l int64, r int64) bool {
		if node == nil {
			return true
		}

		var con bool
		con = int64(node.Val) > l && int64(node.Val) < r

		if !con {
			return false
		}

		var x, y bool
		x = dfs(node.Left, l, int64(node.Val))
		y = dfs(node.Right, int64(node.Val), r)

		return x && y
	}

	var kiri, kanan int64
	kiri = -1000000000000
	kanan = 1000000000000

	return dfs(root, kiri, kanan)
}
