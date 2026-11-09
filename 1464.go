func maxProduct(nums []int) int {
sort.Ints(nums)
k:= (nums[len(nums)-1]-1) * (nums[len(nums)-2]-1)
l:= (nums[0]-1) * (nums[1]-1)
if l>k{
    return l
}else{
    return k
}
}
