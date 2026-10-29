func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	var i, j int
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

	for i = 0; i < n-3; i++ {
		for j = i + 1; j < n-2; j++ {
			var targetBaru int64
			targetBaru = int64(target) - int64(nums[i]) - int64(nums[j])
			
			var kiri, kanan int
			kiri = j + 1
			kanan = n - 1
			
			for kiri < kanan {
				var jumlah int64
				jumlah = int64(nums[kiri]) + int64(nums[kanan])
				if jumlah < targetBaru {
					kiri++
				} else if jumlah > targetBaru {
					kanan--
				} else {
					var kandidat []int
					kandidat = make([]int, 4)
					kandidat[0] = nums[i]
					kandidat[1] = nums[j]
					kandidat[2] = nums[kiri]
					kandidat[3] = nums[kanan]
					hasil = append(hasil, kandidat)
					
					var tempKiri, tempKanan int
					tempKiri = kiri
					tempKanan = kanan
					for kiri < kanan && nums[kiri] == nums[tempKiri] {
						kiri++
					}
					for kiri < kanan && nums[kanan] == nums[tempKanan] {
						kanan--
					}
				}
			}
			
			for j+1 < n && nums[j] == nums[j+1] {
				j++
			}
		}
		for i+1 < n && nums[i] == nums[i+1] {
			i++
		}
	}
	
	return hasil
}
