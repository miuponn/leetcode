func minWindow(s string, t string) string {
    if len(t) > len(s) {                                 // edge case t longer than s, return nil
        return ""
    }
    smap, tmap := make(map[byte]int), make(map[byte]int) // var: maps to keep char -> count for unique chars in t
    need, have := 0, 0                                   // var: desired match count, current match count
    for i := 0; i < len(t); i++ {                        // loop: build count for each char in t
        _, ok := tmap[t[i]]
        if !ok {                                         // if unique char found, increment need
            need++                                        
        }
        tmap[t[i]]++
    }
    for i := 0; i < len(t); i++ {                        // loop: build count for any desired chars in window 
        _, ok := tmap[s[i]]
        if ok {                                           
            smap[s[i]]++
            if smap[s[i]] == tmap[s[i]]{                 // if desired char count satisfied, increment have
                have++
            }
        }
    }
    l, min := 0, math.MaxInt32                           // var: left pointer of window, min res window size so far
    res := []int{-1, -1}                                 // var: indices of min window size 
    if have == need {                                    // if initial window is valid, update res and min 
        res = []int{0, len(t)-1}
        min = len(t)
    }
    for r := len(t); r < len(s); r++ {                   // move window to right for next char of s
        if have < need {                                 // 1. curr window not valid:
            ch := s[r]
            count, ok := tmap[ch]                             
            if ok {                                      //    if new char is desired char, update count in s map
                smap[ch]++
                if smap[ch] == count { have++ }          //    if added desired char satisfies count, increment have
            }
        }
        for have == need {                               // 2. while curr window is valid:
            if (r-l + 1) < min {                         
                min = (r-l + 1)                          //    if curr window is new res min window size so far, 
                res = []int{l, r}                        //    update min, update res indices
            }
            ch := s[l]
            count, ok := tmap[ch]
            if ok {                                      //    move window to left, update count if losing desired char
                smap[ch]--
                if smap[ch] == count-1 { have-- }        //    if count no longer satisfied, decrement have
            }
            l++
        }
    }
    if res[0] == -1 {                                    // edge case no valid windows found in s, return nil
        return ""
    }
    return s[res[0]:res[1]+1]                            // return string within min valid window in s
}