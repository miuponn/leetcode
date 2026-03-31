func largestRectangleArea(heights []int) int {  
    res := 0                                                                  // global max height var
    stack := []int{}                                                          // stack of bar indices

    for i := 0; i < len(heights); i++ {                                       // iterate over bar indices
        for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {   // pop from stack while curr bar shorter than popped bar
            h := heights[stack[len(stack)-1]]                                 // calculate area for popped bar with bar below as 
            stack = stack[:len(stack)-1]                                      // left limit and curr bar as right limit,
            w := i                                                            // width as span between limits and popped bar for height
            if len(stack) != 0 {
                w = i - stack[len(stack)-1] - 1
            }
            res = max(res, h * w)                                             // update global max height var
        }
        stack = append(stack, i)                                              // push curr bar index onto stack
    }
    for len(stack) != 0 {                                                     // for remaining unpopped elements in stack, 
        h := heights[stack[len(stack)-1]]                                     // pop from stack and calculate area,
        n := len(heights)                                                     // using length of heights as right limit 
        stack = stack[:len(stack)-1]                                          // and bar below as left limit
        w := n
        if len(stack) != 0 {                                                  // width as span between limits and popped bar for height
            w = n - stack[len(stack)-1] - 1
        }
        res = max(res, h * w)                                                 // update global max height var
    }
    return res
}
// monotonic increasing stack