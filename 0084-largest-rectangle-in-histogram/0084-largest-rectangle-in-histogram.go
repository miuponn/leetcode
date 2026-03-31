func largestRectangleArea(heights []int) int {  
    res := 0
    stack := []int{}

    for i := 0; i < len(heights); i++ {
        for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            w := i
            if len(stack) != 0 {
                w = i - stack[len(stack)-1] - 1
            }
            res = max(res, h * w)
        }
        stack = append(stack, i)
    }
    for len(stack) != 0 {
        h := heights[stack[len(stack)-1]]
        n := len(heights)
        stack = stack[:len(stack)-1]
        w := n
        if len(stack) != 0 {
            w = n - stack[len(stack)-1] - 1
        }
        res = max(res, h * w)
    }
    return res
}