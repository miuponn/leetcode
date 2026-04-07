func characterReplacement(s string, k int) int {
    m := make(map[byte]int)                 // map var to keep char -> count
    l, maxFreq, res := 0, 0, 0

    for r := 0; r < len(s); r++ {
        m[s[r]]++       // update count in letter
        maxFreq = max(maxFreq, m[s[r]])

        for (r - l + 1) - maxFreq > k {
            m[s[l]]--
            l++
        }

        res = max(res, (r - l + 1))
    }
    return res
}