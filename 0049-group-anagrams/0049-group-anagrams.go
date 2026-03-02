func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)        // init frequencykey=stringsgroup
    for _, str := range strs {                  // iterate over each str in strs
        var key [26]int                         // init array of 26 int vals 
        for _, ch := range str {                // for each ch in str, increment val at alphabet index for key
            key[ch - 'a']++
        }
        groups[key] = append(groups[key], str)  // add str to group for that key
    }
    res := make([][]string, 0, len(groups))     // init slice of subslices with capacity = length of groups
    for _, group := range groups {              // iterate vals in groups and add subslice to slice
        res = append(res, group)
    }
    return res
}
// O(n) = O(n * m + k)