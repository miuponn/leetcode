func isSubsequence(s string, t string) bool {
    i, j := 0, 0
    for i < len(t) && j < len(s) {
        if t[i] == s[j] {
            i, j = i + 1, j+ 1
        } else {
            i++
        }
    }
    if j == len(s) {
        return true
    } else {
        return false
    }
}