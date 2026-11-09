func smallestSubsequence(s string) string {
    lastOccurrence := make([]int, 26)
    for i, ch := range s {
        lastOccurrence[ch-'a'] = i
    }

    inStack := make([]bool, 26)
    stack := make([]byte, 0, len(s))

    for i, ch := range s {
        c := byte(ch)
        if inStack[c-'a'] {
            continue
        }
        for len(stack) > 0 {
            top := stack[len(stack)-1]
            if c < top && lastOccurrence[top-'a'] > i {
                inStack[top-'a'] = false
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        }
        stack = append(stack, c)
        inStack[c-'a'] = true
    }
    return string(stack)
}
