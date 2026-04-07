func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {                          // edge case s1 longer than s2, return false
        return false
    }
    count1, count2 := [26]int{}, [26]int{}          // var: arrays to keep count of 26 lower case letters
    matches := 0                                    // var: match count for each matching char count
    for i := 0; i < len(s1); i++ {                  // loop: add count for first len(s1) chars for s2
        count1[s1[i] - 'a']++
        count2[s2[i] - 'a']++
    }
    for i := 0; i < 26; i++ {                       // loop: count matches for starting window chars
        if count1[i] == count2[i] { matches++ }
    }
    l := 0                                          // left pointer for window
    for r := len(s1); r < len(s2); r++ {            // move window to right for each char in s2
        if matches == 26 {                          // return true if window matches
            return true
        }
        ri := s2[r] - 'a'                           // get index of char at right pointer
        count2[ri]++                                // update count for s2 for char
        if count1[ri] == count2[ri] {               // increment matches if new char count matches
            matches++
        } else if count1[ri] == count2[ri] - 1 {    // decrement matches if new char count unmatched a match
            matches--
        }
        li := s2[l] - 'a'                           // get index of char at left pointer
        count2[li]--                                // update count for s2 for char
        if count1[li] == count2[li] {               // increment matches if new char count matches
            matches++   
        } else if count1[li] == count2[li] + 1 {    // decrement matches if new char count unmatched a match
            matches--
        }
        l++                                         // move left side of window to right 
    }
    return matches == 26                            // return if substring found at end/not found
}