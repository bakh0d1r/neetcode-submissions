func scoreOfParentheses(s string) int {
    res, depth := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            depth++
        } else {
            depth--
            if s[i-1] == '(' {
                res += 1 << depth
            }
        }
    }
    return res
}
