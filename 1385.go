func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var n1, n2, i, j, ans int
	n1 = len(arr1)
	n2 = len(arr2)
	ans = 0

	var a, b, selisih int
	var ada bool

	for i = 0; i < n1; i++ {
		a = arr1[i]
		ada = true

		for j = 0; j < n2; j++ {
			b = arr2[j]
			selisih = a - b
			if selisih < 0 {
				selisih = -selisih
			}

			if selisih <= d {
				ada = false
				break
			}
		}

		if ada {
			ans++
		}
	}

	return ans
}
