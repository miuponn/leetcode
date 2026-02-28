func strStr(haystack string, needle string) int {
    // KMP: LPS array 
    if needle == "" { return 0 }
    lps := make([]int, len(needle))                  // init LPS array
    prevLPS, i := 0, 1                               // pointers for LPS len, array/prefix index
    
    // 1. build list of LPS for each prefix in needle
    for i < len(needle){                             // for prefix in needle,
        if needle[i] == needle[prevLPS] {            // 1. prefix forms new lps with prevLPS
            lps[i] = prevLPS + 1                     //    - add new lps to list
            i++                                      //    - increment next prefix
            prevLPS++                                //    - increment prevLPS 
        } else if prevLPS == 0 {                     // 2. no new lps, no prevLPS
            lps[i] = 0                               //    - add 0 to list
            i++                                      //    - increment next prefix
        } else {                                     // 3. no new lps
            prevLPS = lps[prevLPS - 1]               //    - decrement prevLPS to its prevLPS
        }
    }
    // 2. search for needle in haystack
    i, j := 0, 0                                    // pointers for haystack, needle
    for i < len(haystack){                          // for char i as starting point 
        if haystack[i] == needle[j]{                // 1. char i in haystack == char j in needle
            i++                                     //     - increment char i, j 
            j++
        } else if j == 0 {                          // 2. mismatch, no chars j matched for i
            i++                                     //     - check next char i as starting point
        } else {                                    // 3. mismatch, chars j matched for i
            j = lps[j - 1]                          //     - decrement char j to matched so far
        }
        if j == len(needle) {                       // full match found in haystack
            return i - len(needle)                  //return starting index of needle in haystack
        }
    } 
    return -1
} 

