func isPalindrome(s string) bool {
    i, j := 0, len(s) - 1                                   // two pointers, front and back traversal

    for i < j {                                             // while pointers have not met
        if !(isAlphanumeric(s[i])){ i++ }                   // if ascii, increment/decrement pointer
        if !(isAlphanumeric(s[j])){ j-- }
        if isAlphanumeric(s[i]) && isAlphanumeric(s[j]){    // compare if matching
            if s[i]|32 != s[j]|32 {                         // normalize lower case
                return false
            } 
            i, j = i + 1, j - 1
        }
    }
    return true
}

// helper func: check if byte is alphanumeric
func isAlphanumeric(c byte) bool {
    if c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
        return true
    }
    return false
}