func firstStableIndex(nums []int, k int) int {
    var x, a, max, min, hasil int
    var con bool
    x = len(nums) -1
    con = false
    for i:=0; i<=x; i++{
        a = i
        max = nums[i]
        min = nums[i]
        for j:= 0; j<=a; j++{
            if max < nums[j]{
                max = nums[j]
            }
        }
        for j:= a; j<=x; j++{
            if min > nums[j]{
                min = nums[j]
            }
        }
        fmt.Println(max,min)
        if (max - min) <= k {
            hasil = i
            fmt.Println("break", (max - min), min, max )
            con = true
            break
        }
    }
    if con {
        return hasil
    }else{
        return -1
    }
}
