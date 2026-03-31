func findMin(nums []int) int {
    l, r := 0, len(nums) - 1

    for l < r {
        mid := (l + r) / 2

        if nums[r] < nums[mid] {
            l = mid + 1
        } else if nums[r] > nums[mid] {
            r = mid
        }
    }
    return nums[l]
}