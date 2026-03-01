func canConstruct(ransomNote string, magazine string) bool {
    pool := charPool(magazine) 
    
    // for each char in ransomNote O(m), hashmap lookup
    for i := 0; i < len(ransomNote); i++ {    
        if pool[ransomNote[i]] > 0 {
            pool[ransomNote[i]]--
        } else {
            return false
        }
    }
    return true
}

// build hashmap byte:count for chars in magazine O(n)
func charPool(magazine string) map[byte]int {
    pool := make(map[byte]int)
    for i := 0; i < len(magazine); i++ {
        pool[magazine[i]]++
    }
    return pool
}