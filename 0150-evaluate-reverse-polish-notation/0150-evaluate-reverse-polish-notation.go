func evalRPN(tokens []string) int {
    stack := []int{}                             // init stack 

    for _, token := range tokens {               // iterate over each token
        switch token {
            case "+":                            // if token is operator, pop two operands from stack,
                r := stack[len(stack)-1]         // perform operation and push result back onto stack
                l := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, l+r)
            case "-":
                r := stack[len(stack)-1]
                l := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, l-r)
            case "*":
                r := stack[len(stack)-1]
                l := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, l*r)
            case "/":
                r := stack[len(stack)-1]
                l := stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, l/r)
            default:                                // if token is integer, convert to int and push onto stack
                num, _ := strconv.Atoi(token)
                stack = append(stack, num)
        }
    }
    return stack[0]     // one int element (result) left after iteration complete, pop and return as result
}