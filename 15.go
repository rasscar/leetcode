package main

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

func threeSum(nums []int) [][]int {
	var n, i, j, k, tot int
	n = len(nums)

	urut(nums, 0, n-1)

	var hasil [][]int
	hasil = make([][]int, 0)

	for i = 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j = i + 1
		k = n - 1

		for j < k {
			tot = nums[i] + nums[j] + nums[k]

			if tot > 0 {
				k--
			} else if tot < 0 {
				j++
			} else {
				var temp []int
				temp = make([]int, 3)
				temp[0] = nums[i]
				temp[1] = nums[j]
				temp[2] = nums[k]
				hasil = append(hasil, temp)

				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}

	return hasil
}
