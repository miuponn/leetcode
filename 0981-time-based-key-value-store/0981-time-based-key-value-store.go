type TimeMap struct {
    m map[string][]pair
}

type pair struct {
    timestamp int
    value string
}

func Constructor() TimeMap {
    return TimeMap{m: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.m[key] = append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
    pairs, ok := this.m[key]
    if !ok { return "" }

    lo, hi := 0, len(pairs) - 1
    for lo <= hi {
        mid := (lo + hi)/2
        if pairs[mid].timestamp == timestamp {
            return pairs[mid].value
        } else if pairs[mid].timestamp < timestamp {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    if hi < 0 { return "" } 
    return pairs[hi].value
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */