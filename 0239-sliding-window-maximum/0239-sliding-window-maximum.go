func maxSlidingWindow(nums []int, k int) []int {
    deque, res := []int{}, []int{}                              // vars: deque
    l := 0

    for r := 0; r < len(nums); r++ {                            // for each num in nums, increment right pointer
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[r] {   // dequeue all values less than curr num from right
            deque = deque[:len(deque)-1]                        
        }
        deque = append(deque, r)                                // queue num from right
        if l > deque[0] {                                       // dequeue curr max from left if window shifts out of range
            deque = deque[1:]
        }
        if (r + 1) >= k {                                       // only when curr window is valid:
            res = append(res, nums[deque[0]])                   // add curr max (leftmost of deque) to result
            l++                                                 // shift window
        }
    }
    return res
}