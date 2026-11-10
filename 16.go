func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func urut(ar []int, kiri int, kanan int) {
	if kiri >= kanan {
		return
	}
	var mid int
	mid = (kiri + kanan) / 2
	urut(ar, kiri, mid)
	urut(ar, mid+1, kanan)

	var temp []int
	temp = make([]int, kanan-kiri+1)
	var i, j, k int
	i = kiri
	j = mid + 1
	k = 0

	for i <= mid && j <= kanan {
		if ar[i] <= ar[j] {
			temp[k] = ar[i]
			i++
		} else {
			temp[k] = ar[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = ar[i]
		i++
		k++
	}
	for j <= kanan {
		temp[k] = ar[j]
		j++
		k++
	}
	for i = 0; i < len(temp); i++ {
		ar[kiri+i] = temp[i]
	}
}

func threeSumClosest(nums []int, target int) int {
	var n, i, j, k, tot, terbaik int
	n = len(nums)

	if n < 3 {
		return 0
	}

	terbaik = nums[0] + nums[1] + nums[2]
	urut(nums, 0, n-1)

	for i = 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j = i + 1
		k = n - 1

		for j < k {
			tot = nums[i] + nums[j] + nums[k]

			if tot == target {
				return tot
			}

			if abs(target-tot) < abs(target-terbaik) {
				terbaik = tot
			}

			if tot > target {
				k--
			} else {
				j++
			}
		}
	}

	return terbaik
}
