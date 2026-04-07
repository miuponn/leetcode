func characterReplacement(s string, k int) int {
    m := make(map[byte]int)                         // map var to keep char -> count
    l, maxFreq, res := 0, 0, 0                              

    for r := 0; r < len(s); r++ {                   // increase window by next char in s
        m[s[r]]++                                   // update count for char
        maxFreq = max(maxFreq, m[s[r]])             // update count of most frequent char so far

        for (r - l + 1) - maxFreq > k {             // if chars to replace (window - freq char count) > replaceable
            m[s[l]]--                               // decrement left char, shrink window
            l++                                     
        }

        res = max(res, (r - l + 1))                 // update res with largest valid window so far
    }
    return res
}