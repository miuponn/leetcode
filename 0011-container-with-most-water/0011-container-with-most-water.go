func maxArea(height []int) int {
    i, j := 0, len(height) - 1                          // l, r pointers
    max := 0

    for i < j {                                         // while l < r pointer
        area := (j-i) * min(height[i], height[j])       // calculate area l x w
        if area > max {                                 // update max if area new max
            max = area
        }
        if height[i] <= height[j] {                     // move l if l height shorter
            i++
        } else {                                        // move r if r height shorter
            j--
        }
    }
    return max
}