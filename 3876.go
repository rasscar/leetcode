
func uniformArray(nums1 []int) bool {
    sort.Ints(nums1)

    for p := 0; p <= 1; p++ {
        cnt0, cnt1 := 0, 0
        ok := true

        for _, x := range nums1 {
            par := x & 1

            if par != p {
                need := par ^ p
                if (need == 0 && cnt0 == 0) || (need == 1 && cnt1 == 0) {
                    ok = false
                    break
                }
            }

            if par == 0 {
                cnt0++
            } else {
                cnt1++
            }
        }

        if ok {
            return true
        }
    }

    return false
}
