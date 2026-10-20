func maxDepth(s string) int {
    res := 0
    cur := 0

    for _, c := range s {
        if c == '(' {
            cur++
        } else if c == ')' {
            cur--
        }
        if cur > res {
            res = cur
        }
    }

    return res
}
