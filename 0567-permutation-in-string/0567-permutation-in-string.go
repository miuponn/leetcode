func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    count1, count2 := [26]int{}, [26]int{}
    matches := 0
    for i := 0; i < len(s1); i++ {
        count1[s1[i] - 'a']++
        count2[s2[i] - 'a']++
    }
    for i := 0; i < 26; i++ {
        if count1[i] == count2[i] { matches++ }
    }
    l := 0
    for r := len(s1); r < len(s2); r++ {
        if matches == 26 {
            return true
        }
        ri := s2[r] - 'a'
        count2[ri]++
        if count1[ri] == count2[ri] {
            matches++
        } else if count1[ri] == count2[ri] - 1 {
            matches--
        }
        li := s2[l] - 'a'
        count2[li]--
        if count1[li] == count2[li] {
            matches++
        } else if count1[li] == count2[li] + 1 {
            matches--
        }
        l++
    }
    return matches == 26
}