
func averageOfSubtree(root *TreeNode) int {
	result := 0

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)

		sum := leftSum + rightSum + node.Val
		count := leftCount + rightCount + 1

		if node.Val == sum/count {
			result++
		}

		return sum, count
	}

	dfs(root)
	return result
}

func buildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(values) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func iptr(v int) *int { return &v }
