func smallestIndex(nums []int) int {
    var hasil int
    var con bool
    for i:= 0; i<len(nums);i++{
        var c,has int
        c = nums[i]
        for true{
            var x int
            if c == 0 {
                break
            }
            x = x + (c % 10)
            c = c /10
            has += x
        }
        if has == i {
            hasil = i
            con = true
            break
        }
    }
    if con != true {
        return -1
    }else{
        return hasil
    }
    
}
