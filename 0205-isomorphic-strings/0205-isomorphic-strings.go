func isIsomorphic(s string, t string) bool {
    // map1 for s -> t map2 and t -> s
    map1, map2 := make(map[byte]byte), make(map[byte]byte)

    // for each char in s O(n)
    for i := 0; i < len(s); i++ {
        // check if s !-> t or if t !-> s
        if (map1[s[i]] != t[i] && map1[s[i]] != 0) || (map2[t[i]] != s[i] && map2[t[i]] != 0) {
            return false
        } else {
            map1[s[i]] = t[i]
            map2[t[i]] = s[i]
        }
    }
    return true   
}