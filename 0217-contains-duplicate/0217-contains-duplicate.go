func containsDuplicate(nums []int) bool {
    seen := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if seen[nums[i]] == 1 {
            return true
        } else {
            seen[nums[i]] = 1
        }
    }
    return false
}