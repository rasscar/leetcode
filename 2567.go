func minimizeSum(nums []int) int {
	sort.Ints(nums)
	v := len(nums)

	opsi1 := nums[v-1] - nums[2]
	opsi2 := nums[v-2] - nums[1]
	opsi3 := nums[v-3] - nums[0]

	return mins(opsi1, mins(opsi2, opsi3))
}

func mins(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
