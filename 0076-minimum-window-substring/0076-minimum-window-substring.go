func minWindow(s string, t string) string {
    if len(t) > len(s) {
        return ""
    }
    smap, tmap := make(map[byte]int), make(map[byte]int)
    need, have := 0, 0
    for i := 0; i < len(t); i++ {
        _, ok := tmap[t[i]]
        if !ok {
            need++
        }
        tmap[t[i]]++
    }
    for i := 0; i < len(t); i++ {
        _, ok := tmap[s[i]]
        if ok { 
            smap[s[i]]++
            if smap[s[i]] == tmap[s[i]]{
                have++
            }
        }
    }
    l, min := 0, math.MaxInt32
    res := []int{-1, -1}
    if have == need {
        res = []int{0, len(t)-1}
        min = len(t)
    }
    for r := len(t); r < len(s); r++ {
        if have < need {
            ch := s[r]
            count, ok := tmap[ch]
            if ok {
                smap[ch]++
                if smap[ch] == count { have++ }
            }
        }
        for have == need {
            if (r-l + 1) < min {
                min = (r-l + 1)
                res = []int{l, r}
            }
            ch := s[l]
            count, ok := tmap[ch]
            if ok {
                smap[ch]--
                if smap[ch] == count-1 { have-- }
            }
            l++
        }
    }
    if res[0] == -1 {
        return ""
    }
    return s[res[0]:res[1]+1]
}