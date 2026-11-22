func checkString(s string) bool {
    var a, hit int
    var c string
    stat := true
    a = len(s)
    c = string(s[0])
    hit = 1
        if s[0] == 'b' {
        hit = 0
    }
    for i:=1;i<a;i++{
        if c != string(s[i]){
            if hit == 1{
                c = string(s[i])
                hit --
            }else{
                stat=false
                break
            }
        }
    }
    return stat
}
