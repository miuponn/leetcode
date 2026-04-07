func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2                    // vars A = shorter array, B = other array
    if len(A) > len(B) {
        A, B = B, A
    }
    total := len(A) + len(B)                
    half := (total + 1)/2                   // var for # of elements in left half of merged array
    lo, hi := 0, len(A)                     // pointers for binary searching A for correct partition index
    for lo <= hi {
        i := (lo + hi)/2                    // partition of A = centre of binary search range
        j := half - i                       // partition of B = left half of merged array minus left half of A

        a1 := math.MinInt64                 // var for element left of partition (default: -infinity)
        if i > 0 {                          // if partition not first element, a1 = max of left half
            a1 = A[i - 1]                   
        }
        a2 := math.MaxInt64                 // var for element right of partition (default: infinity)
        if i < len(A){                      // if partition not last element, a2 = min of right half
            a2 = A[i]                       
        }
        b1 := math.MinInt64                 // var for element left of partition (default: -infinity)
        if j > 0 {                          // if partition not first element, b1 = max of left half
            b1 = B[j - 1]                       
        }
        b2 := math.MaxInt64                 // var for element right of partition (default: infinity)
        if j < len(B){                      // if partition not last element, b2 = min of right half
            b2 = B[j]                       
        }
        if a1 <= b2 && b1 <= a2 {           // correct partition indices found for A & B (all left of A,B < all right of A,B)
            if total % 2 == 0 {
                return (float64(max(a1, b1)) + float64(min(a2, b2))) / 2.0 // if even total, return medians/2
            }
            return float64(max(a1, b1))                                  // if odd total, return median
        } else if b1 > a2 {                 // too few elements from A, move partition to right
            lo = i + 1
        } else {                            // too many elements from A, move partition to left
            hi = i - 1          
        }
    }
    return -1
}