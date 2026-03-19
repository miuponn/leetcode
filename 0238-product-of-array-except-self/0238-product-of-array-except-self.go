func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))               // init result slice

    res[0] = 1                                  // prefix prod for first element
    for i := 1; i < len(nums); i++ {            // add prefix prod for all elements
        res[i] = res[i-1] * nums[i-1]
    }
    
    suffix_prod := 1                            // cumulative suffix prod 
    arr_len := len(nums) - 1
    for i := arr_len; i >= 0; i-- {             // iterate right to left
        res[i] = res[i] * suffix_prod           // multiply prefix prod (res[i]) by cumulative suffix prod
        suffix_prod = suffix_prod * nums[i]     // cumulatively build suffix prod right to left
    }
    return res
}

// O(n) time, O(1) space