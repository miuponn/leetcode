func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers) - 1                      // init pointers i,j for first index, last index
    
    for numbers[i] + numbers[j] != target {          // while sum not target
        if numbers[i] + numbers[j] > target { j-- }  // if sum greater than target, decrement j
        if numbers[i] + numbers[j] < target { i++ }  // if sum less than target, increment i
    }
    return []int{i+1, j+1}                           // return indices in array, +1 for 1-indexed 
}
// O(n) time, O(1) space