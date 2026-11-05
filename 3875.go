func uniformArray(nums1 []int) bool {
    odd := 0
    for _, v:= range nums1{
        if v%2==1{
            odd++
        }
    }
    evenPossivle := (odd != 1)
    oddPosible := (odd >=1)
    return evenPossivle || oddPosible
}
