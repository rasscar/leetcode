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

func maxSelectedElements(nums []int) int {
	var n, i, d0, d1, ans, prevVal, curVal, diff int
	n = len(nums)

	if n == 0 {
		return 0
	}

	urut(nums, 0, n-1)

	d0 = 1
	d1 = 1
	ans = 1
	prevVal = nums[0]

	for i = 1; i < n; i++ {
		curVal = nums[i]
		diff = curVal - prevVal
		prevVal = curVal

		if diff == 0 {
			d1 = d0 + 1
		} else if diff == 1 {
			d0 += 1
			d1 += 1
		} else if diff == 2 {
			d0 = d1 + 1
			d1 = 1
		} else {
			d0 = 1
			d1 = 1
		}

		if d0 > ans {
			ans = d0
		}
		if d1 > ans {
			ans = d1
		}
	}

	return ans
}
