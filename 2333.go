func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	diff := make([]int, n)
	k := int64(k1) + int64(k2)
	var total int64
	maxDiff := 0
	for i := 0; i < n; i++ {
		d := nums1[i] - nums2[i]
		if d < 0 {
			d = -d
		}
		diff[i] = d
		total += int64(d)
		if d > maxDiff {
			maxDiff = d
		}
	}
	if total <= k {
		return 0
	}
	kiri := 0
	kanan := maxDiff
	for kiri < kanan {
		mid := kiri + (kanan-kiri)/2
		var operasi int64
		for i := 0; i < n; i++ {
			if diff[i] > mid {
				operasi += int64(diff[i] - mid)
			}
		}
		if operasi <= k {
			kanan = mid
		} else {
			kiri = mid + 1
		}
	}
	batas := kiri
	sisa := k
	for i := 0; i < n; i++ {
		if diff[i] > batas {
			sisa -= int64(diff[i] - batas)
		}
	}
	var hasil int64
	for i := 0; i < n; i++ {
		d := diff[i]
		if d > batas {
			d = batas
		}
		if d == batas && sisa > 0 {
			d--
			sisa--
		}
		hasil += int64(d) * int64(d)
	}
	return hasil
}
