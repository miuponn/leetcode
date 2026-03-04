func longestConsecutive(nums []int) int {
    numsMap := make(map[int]bool)                                    
    longSeq := 0                                                    // var for LCS so far
    
    for _, num := range nums {                                      // pre-process nums to num:bool for O(1) lookup
        numsMap[num] = true
    }

    for num, _ := range numsMap {                                   // iterate over nums array
        if numsMap[num - 1] { continue                              // if not min num, skip sequence processing
        } else {                                                    // if min, count length of sequence
            currSeq := sequence(num, numsMap)
            if currSeq > longSeq { longSeq = currSeq }              // replace LCS if new LCS
        }
    }
    return longSeq  
}

func sequence(num int, map1 map[int]bool) int {                     // helper func for sequence counting
    seq := 1
    for map1[num+seq] {                                             // count length of sequence where num = min
        seq++   
    }
    return seq                                                      // return length of sequence
}