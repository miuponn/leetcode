func searchMatrix(matrix [][]int, target int) bool {
    j := len(matrix[0]) - 1                     // var for last col in row
    lo, hi := 0, len(matrix) - 1               
    for lo <= hi {                              // binary search for row with correct range
        i := (lo + hi) / 2                      // centre row between lo and hi to search
        if target > matrix[i][j] {              // if target greater than last int in centre row, 
            lo = i + 1                          // move lower bound greater than centre row index
        } else if target < matrix[i][0] {       // if target less than first int in centre row,
            hi = i - 1                          // move upper bound less than centre row index
        } else {
            break                               // target within range, centre row found 
        }
    }
    if lo > hi { return false }                 // target in between row ranges, no centre row, early exit
    i := (lo + hi) / 2                          // var for centre row found after first BS
    lo, hi = 0, j                               
    for lo <= hi {                              // binary search for target within row
        j := (lo + hi) / 2                      // centre col between lo and hi to compare
        if target > matrix[i][j] {              // if target greater than centre int,
            lo = j + 1                          // move lower bound greater than centre col index
        } else if target < matrix[i][j] {       // if target less than centre int,
            hi = j - 1                          // move upper bound less than centre col index
        } else {                                // target found, return true
            return true
        }
    }
    return false
}