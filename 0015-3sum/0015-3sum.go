func threeSum(nums []int) [][]int {
    // two sum l + r pointer for last two elements (NSum)
    res := [][]int{}                                                
    sort.Ints(nums)                                             // sort array

    for i := 0; i < len(nums); i++ {                            // for each index i 
        if nums[i] > 0 { break }                                // no solution, skip loop
        if i != 0 && nums[i] == nums[i-1] { continue }          // skip duplicate first element
        l, r := i+1, len(nums)-1                                // init left, right pointers
        for l < r {                                             // while left < right 
            if l != i+1 && nums[l] == nums[l-1] {               // skip duplicate second element
                l++
                continue
            }
            sum := nums[i] + nums[l] + nums[r]                  // calculate sum of 3 elems
            if sum == 0 {                                       // if sum 0, add to result slice, 
                triplet := []int{nums[i], nums[l], nums[r]}     // then move both l and r pointers
                res = append(res, triplet)
                l++                                             
                r--
            } else if sum < 0 {                                 // adjust l pointer if sum too low
                l++
            } else {                                            // adjust r pointer if sum too high
                r--
            }
        }
    }
    return res
}