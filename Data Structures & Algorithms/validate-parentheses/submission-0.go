func isValid(s string) bool {
    brackets := map[byte]byte{
        ')':'(',
        '}':'{',
        ']':'[',
    }
    stack:=[]byte{}
    for i:=range s{
        if len(stack) > 0 && brackets[s[i]] == stack[len(stack)-1] {
            stack=stack[0:len(stack)-1]
            continue
        } 
        stack=append(stack,s[i])
    }
    return len(stack) == 0
}
