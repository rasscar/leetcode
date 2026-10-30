func sortPermutation(nums []int) int {
    k := -1
    for i, x := range nums {
        if x != i {  
            k &= x 
            if k == 0 { return 0 } 
        }
    }
    return max(k, 0) 
}
