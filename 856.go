func scoreOfParentheses(s string) int {
    var y int
    var a,b int
    y = len(s)
    for i:=0;i<y;i++{
        if s[i] == '('{
            a++
        }else{
            a--
            if s[i-1] == '(' {
                b += 1 << a
            }
        }
    }
        return b
}
