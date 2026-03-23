func trap(height []int) int {
    if len(height) == 0 { return 0 }

    l, r := 0, len(height) - 1
    res, lMax, rMax := 0, height[l], height[r]

    for l < r {
        if lMax < rMax {
            l++
            lMax = max(height[l], lMax)
            res += lMax - height[l]
        } else {
            r--
            rMax = max(height[r], rMax)
            res += rMax - height[r]
        }
    }
    return res
}