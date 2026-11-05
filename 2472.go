func cekPalindrome(s string, l, k int) bool{
        for l < k {
                if s[l] != s[k] {
                 return false // bukan palindrome
                }
                l +=1
                k -=1                                                                                         }
        return true // lolos syarat palindrome terkecil
}


func maxPalindromes(s string, k int) int{
var n, count int
n = len(s)
count = 0
        for i:= 0; i+k <= n; {
                if cekPalindrome(s,i,i+k-1){
                        count ++
                        i += k                          // i maju sejauh k
                }else if i+k < n && cekPalindrome(s,i,i+k){
                        i += k+1                        //  memungkinkan untuk sejauh i+k
                        count++
                }else{
                        i++                             // maju biasa jika bukan palindrome alias skip
                }
        }
        return count
}
