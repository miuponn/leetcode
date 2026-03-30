func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))                                  // init res slice with 0 for each temp
    stack := []int{}                                                       // init stack

    for i, temp := range temperatures {                                    // for each item in temps
        // greater temp found -> is next greatest for all elems in stack
        for len(stack) != 0 && temp > temperatures[stack[len(stack)-1]] {  // for all elems currently in stack
            top := stack[len(stack)-1]                                     // get index of top elem to pop
            stack = stack[:len(stack)-1]                                   // pop top elem
            res[top] = i - top                                             // add distance of next greatest for popped elem
        }
        stack = append(stack, i)                                           // push index of curr item onto stack next
    }
    return res
}
// monotonic decreasing stack