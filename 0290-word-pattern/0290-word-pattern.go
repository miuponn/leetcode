func wordPattern(pattern string, s string) bool {
    map1, map2 := make(map[byte]string), make(map[string]byte)           // maps for pattern -> words and words -> pattern
    words := stringSplit(s)

    for i := 0; i < len(pattern); i++ {
        if (map1[pattern[i]] != words[i] && map1[pattern[i]] != "") || (map2[words[i]] != pattern[i] && map2[words[i]] != 0) || len(words) != len(pattern) {
            return false
        } else {
            map1[pattern[i]] = words[i]
            map2[words[i]] = pattern[i]
        }
    }
    return true 
}

// manual string splitter
func stringSplit(s string) []string {
    words := make([]string, 0)                                      // init slice of words
    l := 0                                                          // pointer for first char in word
    for r, ch := range s {                                          
        if ch == ' ' {                                              // if space found, add word l:char to slice
            words = append(words, s[l:r])                           // and shift left pointer to char after space
            l = r+1
        }
    }
    words = append(words, s[l:])                                    // add last word to slice (no trailing spaces)
    return words
}