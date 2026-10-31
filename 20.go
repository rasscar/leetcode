func isValid(s string) bool {
stack := make([]byte, 0, len(s))
for i := 0; i < len(s); i++ {
ch := s[i]
if ch == '(' || ch == '[' || ch == '{' {
stack = append(stack, ch)
} else {
if len(stack) == 0 {
return false
}
top := stack[len(stack)-1]
ada := false
if ch == ')' && top == '(' {
ada = true
}
if ch == ']' && top == '[' {
ada = true
}
if ch == '}' && top == '{' {
ada = true
}
if ada {
stack = stack[:len(stack)-1]
} else {
return false
}
}
}
return len(stack) == 0
}
