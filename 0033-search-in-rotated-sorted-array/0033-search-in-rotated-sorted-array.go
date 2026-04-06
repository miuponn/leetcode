func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1

    for l <= r {
        mid := (l + r) / 2

        if nums[mid] == target {
            return mid
        }
        if nums[l] <= nums[mid] {                               // if pivot in right half or unrotated
            if target < nums[l] || target > nums[mid] {         // if target greater than centre or less than left,
                l = mid + 1                                     // target in right half
            } else {                                            
                r = mid - 1                                     // else, target in left half
            }
        } else {                                                // if pivot in left half
            if target > nums[r] || target < nums[mid] {         // if target less than centre or greater than right,
                r = mid - 1                                     // target in left half
            } else {                                        
                l = mid + 1                                     // else, target in right half
            }
        }
    }
    return -1
}