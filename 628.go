func maximumProduct(nums []int) int {
    sort.Ints(nums)
    x:= len(nums)
    min:= nums[x-1]*nums[0]*nums[1]
    max:= nums[x-1]*nums[x-2]*nums[x-3]
    if min > max {
        return min
    }else{
        return max
    }
}
