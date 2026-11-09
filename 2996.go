func missingInteger(nums []int) int {
    panjang := 0
    for panjang+1 < len(nums) && nums[panjang+1] == nums[panjang]+1 {
        panjang++
    }

    jumlah := 0
    for i := 0; i <= panjang; i++ {
        jumlah += nums[i]
    }

    ada := make(map[int]bool)
    for _, v := range nums {
        ada[v] = true
    }

    x := jumlah
    for ada[x] {
        x++
    }
    return x
}
