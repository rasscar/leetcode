func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

func gcdSum(nums []int) int64 {
	var n, i, mks int
	n = len(nums)

	var pref []int
	pref = make([]int, n)

	mks = nums[0]
	for i = 0; i < n; i++ {
		if nums[i] > mks {
			mks = nums[i]
		}
		pref[i] = gcd(mks, nums[i])
	}

	urut(pref, 0, n-1)

	var l, r int
	l = 0
	r = n - 1

	var ans int64
	ans = 0

	for l < r {
		ans += int64(gcd(pref[l], pref[r]))
		l++
		r--
	}

	return ans
}
