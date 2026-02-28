func isPalindrome(s string) bool {
    i := 0
    j := len(s) - 1

    for i < j {
        if !(isAlphanumeric(s[i])){ i++ }
        if !(isAlphanumeric(s[j])){ j-- }
        if isAlphanumeric(s[i]) && isAlphanumeric(s[j]){
            if s[i]|32 != s[j]|32 {
                return false
            } 
            i, j = i + 1, j - 1
        }
    }
    return true
}

func isAlphanumeric(c byte) bool {
    if c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
        return true
    }
    return false
}