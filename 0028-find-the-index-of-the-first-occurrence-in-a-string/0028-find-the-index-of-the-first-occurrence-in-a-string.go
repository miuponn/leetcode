func strStr(haystack string, needle string) int {
    // outer loop haystack o(n)
    for i := 0; i < len(haystack); i++ {
        // inner loop needle o(m)
        for j := 0; j < len(needle); j++ {
            // compare starting index i + character index j with needle character index j
            if i + j >= len(haystack) || haystack[i+j] != needle[j] {
                // if characters dont match or character out of bounds, break 
                break // next i, j := 0 on break
            }
            if j == len(needle) - 1 { //if full needle found at i, return i
                return i 
            }
        }
    } 
    return -1 //else return not found
}