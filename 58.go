func lengthOfLastWord(s string) int {
    var x int
    var m int
    var a bool
    a = true
    x = len(s)
    if s[x-1] != ' '{
            a=false
        }
    for i :=x-1; i>=0; i--{
        if a!= true{
            if s[i] != ' '{
                m++
            }else{
                break
            }
        }else{
            if s[i] != ' '{
                a = false
                m++
            }
        }
    } 
    return m
}
