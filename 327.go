func countRangeSum(nums []int, lower int, upper int) int {
	var n int
	n = len(nums)
	if n == 0 {
		return 0
	}

	var prefSum []int64
	prefSum = make([]int64, n)
	prefSum[0] = int64(nums[0])

	var i int
	for i = 1; i < n; i++ {
		prefSum[i] = int64(nums[i]) + prefSum[i-1]
	}

	var bawah int64
	bawah = int64(lower)
	var atas int64
	atas = int64(upper)

	var hitungMerge func(kiri int, kanan int) int
	hitungMerge = func(kiri int, kanan int) int {
		if kanan-kiri <= 1 {
			if kanan-kiri == 1 {
				if bawah <= prefSum[kiri] && prefSum[kiri] <= atas {
					return 1
				}
				return 0
			}
			return 0
		}

		var mid int
		mid = (kiri + kanan) / 2
		var jumKiri int
		jumKiri = hitungMerge(kiri, mid)
		var jumKanan int
		jumKanan = hitungMerge(mid, kanan)

		var i, j int
		i = kiri
		j = mid
		var nPtr, mPtr int
		nPtr = 0
		mPtr = 0
		var temp []int64
		temp = make([]int64, kanan-kiri)
		var k int
		k = 0
		var jum int
		jum = 0

		for i < mid {
			for mid+nPtr < kanan && prefSum[mid+nPtr] < prefSum[i]+bawah {
				nPtr++
			}
			for mid+mPtr < kanan && prefSum[mid+mPtr] <= prefSum[i]+atas {
				mPtr++
			}
			for j < kanan && prefSum[j] < prefSum[i] {
				temp[k] = prefSum[j]
				k++
				j++
			}
			temp[k] = prefSum[i]
			k++
			i++
			jum += mPtr - nPtr
		}

		for j < kanan {
			temp[k] = prefSum[j]
			k++
			j++
		}

		var idx int
		for idx = 0; idx < len(temp); idx++ {
			prefSum[kiri+idx] = temp[idx]
		}

		return jumKiri + jumKanan + jum
	}

	return hitungMerge(0, n)
}
