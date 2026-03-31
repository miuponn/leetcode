func findMin(nums []int) int {
    l, r := 0, len(nums) - 1

    for l < r {                          // binary search to converge onto min val
        mid := (l + r) / 2

        if nums[r] < nums[mid] {         // min in later half, converge right
            l = mid + 1
        } else if nums[r] > nums[mid] {  // min in first half, converge left
            r = mid
        }
    }
    return nums[l]
}