func maxSubarrayLength(nums []int, k int) int {
    frek := make(map[int]int)
    kiri := 0
    maks := 0

    for kanan, val := range nums {
        frek[val]++

        for frek[val] > k {
            frek[nums[kiri]]--
            kiri++
        }

        if kanan-kiri+1 > maks {
            maks = kanan - kiri + 1
        }
    }
    return maks
}
