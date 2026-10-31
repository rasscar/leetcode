func removeOuterParentheses(s string) string {
    var answer strings.Builder
    depth := 0

    for i := 0; i < len(s); i++ {
        c := s[i]

        if c == '(' {
            if depth > 0 {
                answer.WriteByte(c)
            }
            depth++
        } else {
            depth--
            if depth > 0 {
                answer.WriteByte(c)
            }
        }
    }

    return answer.String()
}
