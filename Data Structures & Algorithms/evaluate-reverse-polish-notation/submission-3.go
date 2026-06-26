func evalRPN(tokens []string) int {
    add:="+"
    sub:="-"
    mult:="*"
    div:="/"
    stack:=[]int{}
    for i:=range tokens {
        switch tokens[i] {
            case add,sub,div,mult:
                result:=calc(stack,tokens[i])
                stack=stack[:len(stack)-2]
                stack=append(stack,result)
                fmt.Println(stack,tokens[i],result)
            default :
                nm ,_ := strconv.Atoi(tokens[i])
                stack=append(stack,nm)
        }
    }
    return stack[0]
}

func calc(stack []int,operand string) int {
    fmt.Println(stack)
    if len(stack) == 0 {
        return 0
    }
    if len(stack) == 1 {
        return stack[len(stack)-1]
    }
    add:="+"
    sub:="-"
    mult:="*"
    div:="/"
    var last int = stack[len(stack)-1]
    var last2 int = stack[len(stack)-2]
    switch operand {
        case add:
            last2+=last
        case sub:
            last2-=last
        case mult:
            last2*=last
        case div:
            last2/=last
    }
    return last2
}