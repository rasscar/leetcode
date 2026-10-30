func sortableIntegers(nums []int) int {
    var n, hasil int 
    n = len(nums)
    for k:=1;k<=n;k++{
        if n%k == 0 && isPos(nums,k){
            hasil += k
        }
    }
    return hasil
}
func isPos(nums []int,k int) bool {
    maks,maks1,min := -1 << 30, -1 << 30, 1 << 30
    var n, jumlah int
    n = len(nums)
    for i:=0; i<n; i++{
        if (i+1)%k != 0 && nums[i] > nums [1+i]{
            jumlah ++
        }
        if nums[i] < min {
            min = nums[i]
        }
        if nums[i] > maks1 {
            maks1 = nums[i]
        }
        if (i+1)%k  == 0{
            awal := nums[i-k+1]
            akhir := nums[i]
            if jumlah > 1 {
                return false
            
            }
            if jumlah == 1 && akhir > awal{ 
                return false
            }
            if min < maks{
                return false
            }
            maks = maks1
            maks1 = -1 << 30
            min = 1 << 30
            jumlah =0
        }
    }
    return true
}
