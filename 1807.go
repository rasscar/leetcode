func evaluate(s string, knowledge [][]string) string {
    m := make(map[string]string)
    for _, kv := range knowledge {
        m[kv[0]] = kv[1]
    }

    var b strings.Builder

    for i := 0; i < len(s); {
        if s[i] == '(' {
            j := i + 1
            for s[j] != ')' {
                j++
            }
            key := s[i+1 : j]
            if val, ok := m[key]; ok {
                b.WriteString(val)
            } else {
                b.WriteString("?")
            }
            i = j + 1
        } else {
            b.WriteByte(s[i])
            i++
        }
    }

    return b.String()
}
