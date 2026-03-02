func isAnagram(s string, t string) bool {
    map1 := make(map[byte]int)
    count := 0
    for i := 0; i < len(s); i++ {
        if map1[s[i]] == 0 {
            count++
        }
        map1[s[i]]++
    }

    for i := 0; i < len(t); i++ {
        if map1[t[i]] > 0 { 
            map1[t[i]]-- 
        } else {
            return false
        }
        if map1[t[i]] == 0 { count-- }
    }
    return count == 0
}