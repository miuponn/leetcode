func minEatingSpeed(piles []int, h int) int {
    // range of possible k to binary search
    lo, hi := 1, max(piles)
    for lo <= hi {                                          // O(log n) binary search
        // intial k = centre of range 
        k := (lo + hi) / 2 
        hrs := 0
        for _, pile := range piles {                
            hrs += ceil(pile, k)
        }                                               
        if hrs > h {
            lo = k + 1
        } else if hrs <= h {
            hi = k - 1
        }
    }
    return lo
}

// helpers
func ceil(a, b int) int {
    return (a + b - 1) / b
}
// o(n) max func
func max(piles []int) int {
    maximum := 0
    for i := 0; i < len(piles); i++{
        if piles[i] > maximum{
            maximum = piles[i]
        }
    }
    return maximum
}