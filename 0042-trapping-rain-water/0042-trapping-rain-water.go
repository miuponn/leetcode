func trap(height []int) int {
    if len(height) == 0 { return 0 }                // empty list

    l, r := 0, len(height) - 1                      // left, right pointers
    res, lMax, rMax := 0, height[l], height[r]      // left max, right max vars

    for l < r {
        if lMax < rMax {                            // move left if left max < right max
            l++
            lMax = max(height[l], lMax)             // update left max if new max
            res += lMax - height[l]                 // add new water in cusp to total 
        } else {                                    // move right if right max <= left max
            r--                         
            rMax = max(height[r], rMax)             // update right max if new max
            res += rMax - height[r]                 // add new water in cusp to total
        }
    }
    return res
}