func countDaysTogether(s1,s2,s3,s4 string) int {
    var days = [13]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}
    get := func(s string) int{
        a,b,c,d := int(s[0]-'0'), int(s[1]-'0'),int(s[3]-'0'), int(s[4]-'0')
        return days[a*10 + b-1] + c*10 +d
    }
    a,b,c,d := get(s1),get(s2),get(s3),get(s4)
    start, end:= a,b
    if c > start {
        start = c
    }
    if d< end{
        end = d
    }
    if start > end{
        return 0
    }
    return end - start + 1
}
