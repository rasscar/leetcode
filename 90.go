
func subsetsWithDup(nums []int) [][]int {
	var n, i, j int
	n = len(nums)

	for i = 1; i < n; i++ {
		var key int
		key = nums[i]
		j = i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	var hasil [][]int
	hasil = make([][]int, 0)
	var jalur []int
	jalur = make([]int, 0)

	var dfs func(ind int)
	dfs = func(ind int) {
		var salin []int
		salin = make([]int, len(jalur))
		var k int
		for k = 0; k < len(jalur); k++ {
			salin[k] = jalur[k]
		}
		hasil = append(hasil, salin)

		var x int
		for x = ind; x < n; x++ {
			if x > ind && nums[x] == nums[x-1] {
				continue
			}
			jalur = append(jalur, nums[x])
			dfs(x + 1)
			jalur = jalur[:len(jalur)-1]
		}
	}

	dfs(0)
	return hasil
}
