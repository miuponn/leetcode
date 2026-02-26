func lengthOfLastWord(s string) int {
    start := len(s) - 1
    if s[len(s) - 1] == ' '{
        for i := len(s) - 1; i >= 0 && s[i] == ' '; i-- {
            start -= 1
        }
    }
    res := 0
    for i := start; i >= 0 && s[i] != ' '; i-- {
        res += 1
    }
    return res
}