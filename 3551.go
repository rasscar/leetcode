func minSwaps(nums []int) int {
	n:= nums
	length := len(n)
	if length == 0 {
		return 0
	}

	originNums := make([]int, length)
	copy(originNums, n)

	sort.Slice(n, func(i, j int) bool {
		num1 := digitSum(n[i])
		num2 := digitSum(n[j])
		if num1 == num2 {
			return n[i] < n[j]
		}
		return num1 < num2
	})
	sortedMap := map[int]int{}
	for i, num := range n {
		sortedMap[num] = i
	}

	visited := make([]bool, length)
	var dfs func(index int, depth *int) bool
	dfs = func(index int, depth *int) bool {
		if visited[index] {
			return true
		}
		visited[index] = true
		*depth += 1
		return dfs(sortedMap[originNums[index]], depth)
	}

	swapCount := 0
	for i := 0; i < length; i++ {
		if visited[i] {
			continue
		}
		depth := 0
		dfs(i, &depth)
		swapCount += depth - 1
	}

	return swapCount
}

func digitSum(input int) int {
	output := 0
	for input > 0 {
		output += input % 10
		input = input / 10
	}
	return output
}
