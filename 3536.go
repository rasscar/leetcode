func maxProduct(n int) int {
    var has, max2, max, con, next int
    next = n % 10
    con = n / 10 
    n = con / 10
    con = con % 10
    
    if next > con {
        max = next
        max2 = con
    }else{
        max = con
        max2 = next
    }
    for n > 0 {
        next = n % 10
        if (next > max) && (next > max2){
            max2 = max
            max = next
        }else if next >= max2{
            max2 = next

        }
        n = n / 10
    }
    has = max * max2
    return has
}
