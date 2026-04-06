type TimeMap struct {                                                   // store map of key -> slice of (timestamp, value)
    m map[string][]pair
}

type pair struct {                                                      // helper struct pair for (timestamp, value)
    timestamp int
    value string
}

func Constructor() TimeMap {                                            // init m map with key -> slice of (timestamp, value)
    return TimeMap{m: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {    // add (timestamp, value) to key -> slice
    this.m[key] = append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {        
    pairs, ok := this.m[key]                                            // return nil of key doesnt exist in map
    if !ok { return "" }

    lo, hi := 0, len(pairs) - 1                                         
    for lo <= hi {                                                      // binary search for target in map
        mid := (lo + hi)/2
        if pairs[mid].timestamp == timestamp {
            return pairs[mid].value                                     // return value at target if target in map
        } else if pairs[mid].timestamp < timestamp {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    if hi < 0 { return "" }                                            // if target not in map and target < all timestamps, return nil
    return pairs[hi].value                                             // if target not in map, return rightmost timestamp less than target
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */