func topKFrequent(nums []int, k int) []int {
    frequency := make(map[int]int)                      // init map for keeping num:count

    for _, num := range nums{                           // iterate over nums and populate map
        frequency[num]++
    }
    buckets := make([][]int, len(nums) + 1)             // init slice of slices (buckets) where index = frequency

    for num, count := range frequency {                 // iterate over each num in frequency map
        buckets[count] = append(buckets[count], num)    // append to correct bucket (frequency = index)
    }

    res := []int{}                                      // init results slice
    for i := len(nums); i > 0 && len(res) != k; i--{    // iterate down i from most frequent (len of nums) until k
        res = append(res, buckets[i]...)                // add nums from buckets for frequency i
    }
    return res
}