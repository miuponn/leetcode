func search(nums []int, target int) int {
    lo, hi := 0, len(nums) - 1              // lower, upper bound incides span sorted array

    for lo <= hi {
        mid := (lo + hi) / 2                // centre index

        if target > nums[mid] {             // if target greater than centre value
            lo = mid + 1                    // move lower bound greater than centre index
        } else if target < nums[mid] {      // if target less than centre value
            hi = mid - 1                    // move upper bound less than centre index
        } else {
            return mid                      // target found at centre value, return centre index
        }
    }
    return -1
}